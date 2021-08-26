<template>
  <div>
    <v-row>
      <v-col>
        <v-btn rounded color="primary">
          <v-icon left>mdi-plus</v-icon>
          {{ $t("add_backlog") }}
        </v-btn>
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <div v-show="showModalVar">
          <div id="app">
            <v-app id="inspire">
              <div class="text-center"></div>
            </v-app>
          </div>
        </div>
        <v-card class="mx-auto rounded-xl">
          <template v-if="!loadingStories && stories.length > 0">
            <v-simple-table>
              <template v-slot:default>
                <thead>
                  <tr>
                    <th class="text-left">{{ $t("#") }}</th>
                    <th class="text-left">{{ $t("as") }}</th>
                    <th class="text-left">{{ $t("because") }}</th>
                    <th class="text-left">{{ $t("want") }}</th>
                    <th></th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="story in stories" :key="story.ID">
                    <td>U{{ zeroPad(story.Number, 3) }}</td>
                    <td>{{ story.As }}</td>
                    <td>{{ story.Because }}</td>
                    <td>{{ story.Want }}</td>
                    <td class="text-center">
                      <v-btn
                        id="show-modal"
                        @click="showUserStoryModal(story)"
                        icon
                      >
                        <v-icon>mdi-dots-horizontal-circle-outline</v-icon>
                      </v-btn>
                    </td>
                  </tr>
                </tbody>
              </template>
            </v-simple-table>
          </template>
          <template v-if="!loadingStories && stories.length == 0">
            <v-card-text class="text-center">
              <v-img src="/no_result.png" max-height="64" contain />
              No variance energy.
            </v-card-text>
          </template>
          <template v-if="loadingStories">
            <div class="text-center">
              <v-progress-circular
                indeterminate
                color="primary"
                size="20"
              ></v-progress-circular>
            </div>
          </template>
        </v-card>
      </v-col>
    </v-row>

    <v-dialog v-model="dialog" width="90%">
      <v-card style="border-radius:30px">
        <v-card-title class="text-h6 grey lighten-2">
          Subtasks
          <v-card-actions class="text-right">
            <v-btn icon @click="dialog = false">
              <v-icon>mdi-minus</v-icon>
            </v-btn>
          </v-card-actions>
        </v-card-title>
        <v-simple-table style="border-radius:30px">
          <template v-slot:default >
            <thead style="width: 100%">
              <tr>
                <th style="width: 5%">{{ $t("#") }}</th>
                <th style="width: 30%">{{ $t("text") }}</th>
                <th style="width: 40%">{{ $t("status") }}</th>
                <th style="width: 20%">{{ $t("point") }}</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="subtask in subtasks" :key="subtask.ID">
                <td>{{ subtask.ID }}</td>
                <td>{{ subtask.Text }}</td>
                <td>
                  <v-col class="d-flex" cols="12" sm="6">
                    <v-select :items="items" dense solo style="height: 40px">
                      <template #item="{ item }">
                          <span
                            :class="[
                              { success: item.includes('yellow') },
                              { info: item.includes('blue') },
                              { warning: item.includes('grey') },
                            ]"
                          >
                            {{ item }}</span
                          >
                      </template></v-select
                    >
                  </v-col>
                </td>

                <td>{{ subtask.Point }}</td>
              </tr>
            </tbody>
          </template>
        </v-simple-table>

        <v-divider></v-divider>
      </v-card>
    </v-dialog>
  </div>
</template>
<script>
export default {
  data() {
    return {
      items: ["yellow", "blue", "grey", "Buzz"],
      checkbox: false,
      dialog: false,
      showModalVar: false,
      stories: [],
      loadingStories: false,
      subtasks: [],
    };
  },
  mounted() {
    this.loadStories();
  },
  methods: {
    async showUserStoryModal(story) {
      this.dialog = true;
      this.subtasks = await this.$axios.$get(
        `/api/projects/${this.$route.params.id}/stories/${story.ID}/subtasks`
      );
    },
    async loadStories() {
      this.loadingStories = true;
      this.stories = await this.$axios.$get(
        `/api/projects/${this.$route.params.id}/stories`
      );
      this.loadingStories = false;
    },
    zeroPad(num, places) {
      var zero = places - num.toString().length + 1;
      return Array(+(zero > 0 && zero)).join("0") + num;
    },
  },
};
</script>
<style scoped>
.pink {
  background-color: yellow;
}
</style>