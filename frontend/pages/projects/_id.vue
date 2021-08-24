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
                      <v-btn icon>
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
  </div>
</template>
<script>
export default {
  data() {
    return {
      stories: [],
      loadingStories: false,
    };
  },
  mounted() {
    this.loadStories();
  },
  methods: {
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
