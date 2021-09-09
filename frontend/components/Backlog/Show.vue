<template>
  <v-card>
    <v-card-title class="mt-4 mb-2 flex-column">
      <div
        class="
          text-headline
          primary--text
          mb-2
          d-flex
          align-center
          justify-space-between
        "
        style="width: 100%"
      >
        <div class="ml-3">
          {{ backlog.Number > 0 ? $backlogNumber(backlog) : "####" }}
        </div>
        <v-select
          solo
          flat
          hide-details
          style="max-width: 200px"
          class="ml-2"
          :items="types"
          :label="$t('type')"
          item-text="name"
          item-value="value"
          v-model="backlog.Type"
          name="save_backlog_type"
          @change="updateBacklog({ Type: backlog.Type })"
        ></v-select>
      </div>
      <textarea
        style="width: 100%; word-break: normal"
        class="text-h5 text-center font-weight-light"
        v-model="backlog.Text"
        ref="taskarea"
        @change="updateBacklog({ Text: backlog.Text })"
      ></textarea>
      <div class="mt-4 mb-4" style="width: 50%">
        <v-divider></v-divider>
      </div>
      <div class="text-center mb-1">
        <v-btn-toggle
          v-model="backlog.Status"
          borderless
          mandatory
          @change="updateBacklog({ Status: backlog.Status })"
        >
          <v-btn
            v-for="status in $statuses.backlog"
            class="rounded-pill"
            small
            :key="status"
            :value="status"
            :color="
              backlog.Status == status
                ? `${$colors.status[status]}`
                : 'transparent'
            "
            :class="
              backlog.Status == status && backlog.Status != 'todo'
                ? 'white--text'
                : ''
            "
          >
            {{ $t(status) }}
          </v-btn>
        </v-btn-toggle>
      </div>
      <div class="mt-4 mb-4" style="width: 50%">
        <v-divider></v-divider>
      </div>
      <!-- <div class="text-center">
        <v-btn-toggle
          v-model="backlog.Score"
          color="primary"
          borderless
          mandatory
        >
          <v-btn
            v-for="score in $scores"
            class="rounded-pill"
            :key="score"
            :value="score"
            small
          >
            {{ score }}
          </v-btn>
        </v-btn-toggle>
      </div>
      <div class="mt-4 mb-4" style="width: 50%">
        <v-divider></v-divider>
      </div> -->
    </v-card-title>
    <v-card-text>
      <v-card outlined>
        <v-row v-for="subtask in subtasks" :key="subtask.ID" no-gutters dense>
          <v-col cols="9">
            <v-textarea
              solo
              flat
              class="pl-2 my-1"
              v-model="subtask.Text"
              rows="1"
              auto-grow
              hide-details
              no-resize
            ></v-textarea>
            <v-divider></v-divider>
          </v-col>
          <v-col cols="1">
            <div
              class="d-flex justify-center align-center"
              style="height: 100%"
            >
              {{ subtask.Point }}
            </div>
            <v-divider></v-divider>
          </v-col>
          <v-col cols="2" style="min-width: 120px" class="p-0">
            <div
              class="d-flex justify-center align-center"
              style="height: 100%"
            >
              <Status :status="subtask.Status" type="subtask" small></Status>
            </div>
            <v-divider></v-divider>
          </v-col>
        </v-row>
      </v-card>
    </v-card-text>
  </v-card>
</template>
<script>
export default {
  props: ["backlog"],
  data() {
    return {
      loadingSubtasks: false,
      subtasks: [],
      types: [
        {
          name:
            this.$backlogEmoji("user_story") + "   " + this.$t("user_story"),
          value: "user_story",
        },
        {
          name:
            this.$backlogEmoji("technical_task") +
            "   " +
            this.$t("technical_task"),
          value: "technical_task",
        },
        {
          name: this.$backlogEmoji("bug") + "   " + this.$t("bug"),
          value: "bug",
        },
      ],
    };
  },
  watch: {
    backlog: {
      handler: function () {
        var self = this;
        this.loadSubtasks();
        if (this.backlog.ID == 0) {
          setTimeout(() => self.$refs.taskarea.focus(), 100);
        }
      },
      immediate: true,
    },
  },
  methods: {
    async loadSubtasks() {
      if (this.backlog.ID == 0) {
        return;
      }
      this.loadingSubtasks = true;
      this.subtasks = await this.$axios.$get(
        `/api/projects/${this.$route.params.id}/backlogs/${this.backlog.ID}/subtasks`
      );
      this.loadingSubtasks = false;
    },
    async saveBacklog() {
      if (this.backlog.ID > 0) {
        return;
      }
      const data = await this.$axios.$post(
        `/api/projects/${this.$route.params.id}/backlogs`,
        this.backlog
      );
      this.$emit("update:backlog", data);
      this.$nuxt.$emit("added-backlog", this.backlog);
    },
    async updateBacklog(data) {
      await this.saveBacklog();
      await this.$axios.$put(
        `/api/projects/${this.$route.params.id}/backlogs/${this.backlog.ID}`,
        data
      );
      this.$emit("update:backlog", this.backlog);
    },
  },
};
</script>