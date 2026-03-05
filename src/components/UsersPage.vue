<template>
  <v-container fluid class="pa-6">
    <v-card elevation="2" rounded="lg">
      <v-card-title class="d-flex align-center bg-primary text-white py-4 px-6">
        <v-icon icon="mdi-account-group" class="mr-3"></v-icon>
        <span class="text-h5 font-weight-bold">User Management</span>
        <v-spacer></v-spacer>
        <v-btn
          color="white"
          variant="elevated"
          prepend-icon="mdi-account-plus"
          class="text-primary font-weight-bold"
          @click="dialog = true"
        >
          Add User
        </v-btn>
      </v-card-title>

      <v-card-text class="pa-4">
        <v-text-field
          v-model="search"
          label="Search Users"
          prepend-inner-icon="mdi-magnify"
          variant="outlined"
          density="comfortable"
          hide-details
          class="mb-4"
        ></v-text-field>

        <v-data-table
          :headers="headers"
          :items="usersList"
          :search="search"
          class="elevation-0"
          hover
        >
          <template v-slot:[`item.active`]="{ item }">
            <div class="d-flex align-center justify-center">
              <v-btn
                :color="item.active ? 'error' : 'success'"
                :variant="item.active ? 'tonal' : 'elevated'"
                size="small"
                class="font-weight-bold px-4"
                rounded="pill"
                :loading="togglingId === item.userid"
                @click="handleToggle(item)"
                min-width="120"
              >
                <v-icon
                  :icon="item.active ? 'mdi-account-off' : 'mdi-account-check'"
                  start
                  size="small"
                ></v-icon>
                {{ item.active ? "Deactivate" : "Activate" }}
              </v-btn>
              <v-tooltip activator="parent" location="top">
                {{
                  item.active
                    ? "Click to deactivate user"
                    : "Click to activate user"
                }}
              </v-tooltip>
            </div>
          </template>
          <template v-slot:[`item.role`]="{ item }">
            <v-chip
              :color="getRoleColor(item.role)"
              size="small"
              class="text-uppercase font-weight-bold"
              label
            >
              {{ item.role }}
            </v-chip>
          </template>
        </v-data-table>
      </v-card-text>
    </v-card>

    <!-- Add User Dialog -->
    <v-dialog v-model="dialog" max-width="500px">
      <v-card rounded="lg">
        <v-card-title class="bg-primary text-white py-4 px-6">
          <span class="text-h6">Add New User</span>
        </v-card-title>

        <v-card-text class="pa-6">
          <v-form ref="form" v-model="valid">
            <v-text-field
              v-model="newUser.userid"
              label="User ID"
              variant="outlined"
              density="comfortable"
              :rules="[(v) => !!v || 'User ID is required']"
              required
            ></v-text-field>
            <v-text-field
              v-model="newUser.password"
              label="Password"
              type="password"
              variant="outlined"
              density="comfortable"
              :rules="[(v) => !!v || 'Password is required']"
              required
            ></v-text-field>
            <v-text-field
              v-model="newUser.name"
              label="Full Name"
              variant="outlined"
              density="comfortable"
              :rules="[(v) => !!v || 'Name is required']"
              required
            ></v-text-field>
            <v-text-field
              v-model="newUser.mobile"
              label="Mobile Number"
              variant="outlined"
              density="comfortable"
              type="tel"
              :rules="[
                (v) => !!v || 'Mobile is required',
                (v) => /^[0-9]{10}$/.test(v) || 'Mobile must be 10 digits',
              ]"
              required
            ></v-text-field>
            <v-text-field
              v-model="newUser.email"
              label="Email"
              variant="outlined"
              density="comfortable"
              type="email"
              :rules="[
                (v) => !!v || 'Email is required',
                (v) => /.+@.+\..+/.test(v) || 'E-mail must be valid',
              ]"
              required
            ></v-text-field>
            <v-select
              v-model="newUser.role"
              :items="['admin', 'manager', 'biller']"
              label="Role"
              variant="outlined"
              density="comfortable"
              :rules="[(v) => !!v || 'Role is required']"
              required
            ></v-select>
          </v-form>
        </v-card-text>

        <v-card-actions class="pa-4 pt-0">
          <v-spacer></v-spacer>
          <v-btn color="grey-darken-1" variant="text" @click="closeDialog">
            Cancel
          </v-btn>
          <v-btn
            color="primary"
            variant="elevated"
            :disabled="!valid"
            :loading="loading"
            @click="saveUser"
          >
            Save User
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <v-snackbar v-model="snackbar" :color="snackbarColor" timeout="3000">
      {{ snackbarMessage }}
    </v-snackbar>
  </v-container>
</template>

<script>
import EventService from "../services/eventServices";

export default {
  name: "UsersPage",
  data() {
    return {
      search: "",
      dialog: false,
      valid: false,
      loading: false,
      snackbar: false,
      headers: [
        { title: "User ID", key: "userid" },
        { title: "Name", key: "name" },
        { title: "Role", key: "role", align: "center" },
        {
          title: "Status Action",
          key: "active",
          align: "center",
          sortable: false,
        },
      ],
      snackbarMessage: "",
      snackbarColor: "success",
      togglingId: null,
      usersList: [], // Local state for users
      newUser: {
        userid: "",
        password: "",
        name: "",
        mobile: "",
        email: "",
        role: "biller",
      },
    };
  },
  computed: {
    // Removed dependency on store for users
  },
  mounted() {
    this.fetchUsers();
  },
  methods: {
    async fetchUsers() {
      this.loading = true;
      try {
        const response = await EventService.getUsers();
        // Backend returns array of users. Map them if necessary or use directly
        // The backend returns: user_id, user_name, user_role, account_status
        this.usersList = response.data.map((u) => ({
          userid: u.user_id,
          name: u.user_name,
          role: u.user_role,
          active: u.account_status === "ACTIVE",
        }));
      } catch (error) {
        console.error("Error fetching users:", error);
        this.showSnackbar("Error fetching users", "error");
      } finally {
        this.loading = false;
      }
    },
    getRoleColor(role) {
      switch (role) {
        case "admin":
          return "error";
        case "manager":
          return "primary";
        case "biller":
          return "info";
        default:
          return "grey";
      }
    },
    closeDialog() {
      this.dialog = false;
      this.$refs.form.reset();
      this.newUser = {
        userid: "",
        password: "",
        name: "",
        mobile: "",
        email: "",
        role: "biller",
      };
    },
    async saveUser() {
      if (!this.valid) return;
      this.loading = true;
      try {
        const payload = {
          user_id: this.newUser.userid,
          password: this.newUser.password,
          user_name: this.newUser.name,
          user_role: this.newUser.role,
          user_email: this.newUser.email,
          user_phone: this.newUser.mobile,
          created_by: "Admin", // TODO: Replace with actual logged-in user
        };
        await EventService.createUser(payload);
        this.snackbarMessage = "User added successfully!";
        this.snackbarColor = "success";
        this.snackbar = true;
        this.closeDialog();
        this.fetchUsers(); // Refresh list
      } catch (error) {
        console.error("Error saving user:", error);
        this.snackbarMessage = "Error saving user";
        this.snackbarColor = "error";
        this.snackbar = true;
      } finally {
        this.loading = false;
      }
    },
    async handleToggle(user) {
      if (this.togglingId) return;
      this.togglingId = user.userid;
      try {
        const newStatus = !user.active;
        await EventService.updateUserStatus(user.userid, newStatus);

        // Update local state
        user.active = newStatus;

        this.snackbarMessage = `User ${user.userid} ${
          !user.active ? "deactivated" : "activated"
        }`;
        this.snackbarColor = user.active ? "success" : "warning";
        this.snackbar = true;
      } catch (error) {
        console.error("Error toggling user status:", error);
        this.snackbarMessage = "Error updating user status";
        this.snackbarColor = "error";
        this.snackbar = true;
      } finally {
        this.togglingId = null;
      }
    },
    showSnackbar(message, color) {
      this.snackbarMessage = message;
      this.snackbarColor = color;
      this.snackbar = true;
    },
  },
};
</script>
