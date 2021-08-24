<template>
  <v-navigation-drawer
    app
    :permanent="$vuetify.breakpoint.mdAndUp"
    hide-overlay
    v-model="show"
    width="240px"
  >
    <v-list-item>
      <v-list-item-content>
        <v-list-item-title class="text-h6"> OUStory </v-list-item-title>
        <v-list-item-subtitle> Sprint scrum management </v-list-item-subtitle>
      </v-list-item-content>
    </v-list-item>
    <v-divider></v-divider>
    <v-list nav dense class="grow" rounded>
      <div v-for="item in items" :key="item.name">
        <v-list-item-group color="primary darken-1">
          <v-subheader class="text-uppercase mt-2" v-if="item.subheader">
            {{ item.subheader }}
          </v-subheader>
          <v-list-item v-else :to="localePath(item.link)" class="mt-1">
            <v-list-item-icon v-if="item.icon">
              <v-icon>{{ item.icon }}</v-icon>
            </v-list-item-icon>
            <v-list-item-avatar v-if="item.img" size="24">
              <v-img :src="item.img"></v-img>
            </v-list-item-avatar>
            <v-list-item-content>
              <v-list-item-title>{{ item.name }}</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </v-list-item-group>
      </div>
      <v-subheader class="text-uppercase mt-2">
        {{ $t("projects") }}
      </v-subheader>
      <template v-if="!loadingProjects && projects.length > 0">
        <v-list-item-group color="primary darken-1">
          <v-list-item
            class="mt-1"
            v-for="project in projects"
            :key="project.ID"
            :to="`/projects/${project.ID}`"
          >
            <v-list-item-avatar v-if="project.Logo" size="24">
              <v-img :src="project.Logo"></v-img>
            </v-list-item-avatar>
            <v-list-item-content>
              <v-list-item-title>{{ project.Name }}</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </v-list-item-group>
      </template>
      <template v-if="!loadingProjects && projects.length == 0">
        <div class="text-center text--secondary">
          <v-icon>mdi-alien-outline</v-icon>
          <br />
          {{ $t("no_projects_exist") }}
        </div>
      </template>
      <template v-if="loadingProjects">
        <div class="text-center">
          <v-progress-circular
            indeterminate
            color="primary"
            size="20"
          ></v-progress-circular>
        </div>
      </template>
    </v-list>
  </v-navigation-drawer>
</template>

<script>
export default {
  props: ["drawer"],
  model: {
    prop: "drawer",
    event: "drawerChange",
  },
  watch: {
    drawer: function (val) {
      this.show = val;
    },
    show: function (val) {
      if (!val) {
        this.$emit("update:drawer", val);
      }
    },
  },
  data() {
    return {
      show: this.drawer,
      items: [
        {
          icon: "mdi-home",
          name: this.$t("home"),
          link: "/",
        },
      ],
      projects: [],
      loadingProjects: true,
    };
  },
  mounted() {
    this.loadProjects();
  },
  methods: {
    async loadProjects() {
      this.loadingProjects = true;
      this.projects = await this.$axios.$get("/api/projects");
      this.loadingProjects = false;
    },
  },
};
</script>