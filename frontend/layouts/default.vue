<template>
  <v-app dark>
    <NavigationTop :drawer.sync="drawer"></NavigationTop>
    <NavigationDrawer :drawer.sync="drawer"></NavigationDrawer>
    <v-main style="background-color: #f0f0f0">
      <v-container>
        <Nuxt />
      </v-container>
    </v-main>
    <v-footer absolute app>
      <span>Made with &#10084;&#65039; in {{ new Date().getFullYear() }}</span>
      <v-spacer />
      <span>Version 1.0</span>
    </v-footer>
  </v-app>
</template>

<script>
export default {
  data() {
    return {
      drawer: false,
    };
  },
};
</script>
