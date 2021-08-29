<template>
  <v-btn
    style="font-size: 0.8rem; height: 100%; width: 100%"
    class="d-flex justify-center align-center text-uppercase font-weight-medium"
    depressed
    tile
    :class="$colors.status[status]"
    @click="changeStatus()"
  >
    {{ $t(status) }}
  </v-btn>
</template>

<script>
export default {
  props: ["status", "type"],
  methods: {
    changeStatus() {
      console.log("Clicked");
    },
  },
};
</script>