<template>
  <v-app-bar
    app
    flat
    style="background-color: white"
  >
    <v-app-bar-nav-icon
      @click="toggleDrawer()"
      v-if="!$vuetify.breakpoint.mdAndUp"
    ></v-app-bar-nav-icon>

    {{$t('put_something_here')}}
  </v-app-bar>
</template>

<script>
export default {
  props: ["drawer"],
  model: {
    prop: "drawer",
    event: "drawerChange",
  },
  data() {
    return {};
  },
  methods: {
    toggleDrawer() {
      this.$emit("update:drawer", !this.drawer);
    },
  },
};
</script>

