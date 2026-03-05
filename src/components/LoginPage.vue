<template>
  <v-container
    fluid
    class="fill-height bg-grey-lighten-3 d-flex align-center justify-center"
  >
    <v-card elevation="8" width="100%" max-width="400" rounded="lg">
      <v-card-item class="bg-primary pt-6 pb-4 text-center">
        <div class="d-flex justify-center mb-2 mt-4">
          <TestLogo />
        </div>
        <div class="text-h5 font-weight-bold text-white">Secure Login</div>
        <div class="text-subtitle-2 text-white opacity-80">
          Enter your credentials to continue
        </div>
      </v-card-item>

      <v-card-text class="pa-6 pt-8">
        <v-form ref="form" @submit.prevent="login">
          <v-text-field
            v-model="userid"
            label="User ID"
            prepend-inner-icon="mdi-account"
            variant="outlined"
            density="comfortable"
            color="primary"
            class="mb-4"
            :rules="[(v) => !!v || 'User ID is required']"
          ></v-text-field>

          <v-text-field
            v-model="password"
            label="Password"
            prepend-inner-icon="mdi-lock"
            :append-inner-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
            :type="showPassword ? 'text' : 'password'"
            @click:append-inner="showPassword = !showPassword"
            variant="outlined"
            density="comfortable"
            color="primary"
            class="mb-6"
            :rules="[(v) => !!v || 'Password is required']"
          ></v-text-field>

          <v-btn
            block
            color="primary"
            size="large"
            type="submit"
            :loading="loading"
            class="font-weight-bold text-uppercase"
            elevation="2"
          >
            Login
          </v-btn>
        </v-form>
      </v-card-text>
    </v-card>

    <v-snackbar
      v-model="snackbar.show"
      :color="snackbar.color"
      location="top"
      timeout="3000"
    >
      <div class="d-flex align-center">
        <v-icon :icon="snackbar.icon" class="mr-2"></v-icon>
        {{ snackbar.message }}
      </div>
    </v-snackbar>
  </v-container>
</template>

<script>
import TestLogo from "@/components/TestLogo.vue";
import EventService from "@/services/eventServices";

export default {
  name: "LoginPage",
  components: {
    TestLogo,
  },
  data() {
    return {
      userid: "",
      password: "",
      showPassword: false,
      loading: false,
      snackbar: {
        show: false,
        message: "",
        color: "success",
        icon: "mdi-check-circle",
      },
    };
  },
  methods: {
    async login() {
      const { valid } = await this.$refs.form.validate();
      if (!valid) return;

      this.loading = true;

      try {
        const response = await EventService.login({
          user_id: this.userid,
          password: this.password,
        });

        if (response.data.status) {
          const backendUser = response.data.details;

          // Map backend response to store's expected format
          const userDetails = {
            userid: this.userid, // backend doesn't return userid in details struct, so use input
            name: backendUser.user_name,
            role: backendUser.user_role, // Store expects 'role'
            mobile: backendUser.user_mobile,
            email: backendUser.user_email,
          };

          console.log("Backend User:", backendUser);
          console.log("Mapped User Details:", userDetails);

          localStorage.setItem("user", JSON.stringify(userDetails));
          localStorage.setItem("isAuthenticated", "true");

          // Dispatch to store
          this.$store.commit("SET_USER", userDetails);

          this.showSnackbar(
            "Login Successful! Redirecting...",
            "success",
            "mdi-check-circle"
          );

          setTimeout(() => {
            const role = userDetails.role
              ? userDetails.role.toLowerCase().trim()
              : "";
            console.log("Processed Role for Redirect:", role);

            if (role === "biller") {
              this.$router.push({ name: "billing" });
            } else {
              this.$router.push({ name: "dashboard-home" });
            }
          }, 1000);
        } else {
          throw new Error(response.data.message || "Login failed");
        }
      } catch (error) {
        console.error("Login error:", error);
        this.showSnackbar(
          error.message || "Invalid Credentials",
          "error",
          "mdi-alert-circle"
        );
      } finally {
        this.loading = false;
      }
    },
    showSnackbar(message, color, icon) {
      this.snackbar.message = message;
      this.snackbar.color = color;
      this.snackbar.icon = icon;
      this.snackbar.show = true;
    },
  },
};
</script>
