<template>
  <v-btn
    :style="`${isSmall ? yesSmall.style : notSmall.style}`"
    :class="`${isSmall ? yesSmall.class : notSmall.class}  ${
      status != 'todo' ? 'white--text' : ''
    }`"
    depressed
    tile
    :small="isSmall"
    @click="changeStatus()"
    :color="$colors.status[status]"
  >
    {{ $t(status) }}
  </v-btn>
</template>

<script>
export default {
  props: ["status", "type", "small"],
  methods: {
    changeStatus() {
      console.log(this.small);
    },
  },
  computed: {
    isSmall: function () {
      return this.small == "";
    },
  },
  data() {
    return {
      notSmall: {
        style: "font-size: 0.8rem; height: 100%; width: 100%",
        class:
          "d-flex justify-center align-center text-uppercase font-weight-medium",
      },
      yesSmall: {
        style: "",
        class:
          "d-flex justify-center align-center text-uppercase font-weight-medium rounded-pill",
      },
    };
  },
};
</script>