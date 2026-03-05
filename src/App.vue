<template>
  <v-app>
    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>

<script>
export default {
  name: "App",

  data: () => ({
    //
  }),
  created() {
    const user = localStorage.getItem("user");
    if (user) {
      this.$store.commit("SET_USER", JSON.parse(user));
    }
  },
};
</script>
