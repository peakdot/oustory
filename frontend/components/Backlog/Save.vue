<template>
  <v-form v-model="valid" ref="form" @submit.prevent="saveBacklog">
    <v-card>
      <v-card-title class="primary--text">
        {{ $t("add_backlog") }}
      </v-card-title>
      <v-card-text>
        <v-select
          :items="types"
          :label="$t('type')"
          item-text="name"
          item-value="value"
          v-model="backlog.Type"
          name="save_backlog_type"
        ></v-select>
        <v-textarea
          :label="$t('task')"
          name="save_backlog_text"
          v-model="backlog.Text"
          :rules="taskRules"
          autofocus
          auto-grow
        ></v-textarea>
        <div class="text-center">
          <v-btn-toggle v-model="backlog.Status" borderless mandatory>
            <v-btn
              v-for="status in $statuses.backlog"
              class="rounded-pill"
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
      </v-card-text>
      <v-card-actions>
        <v-btn
          rounded
          block
          x-large
          color="primary"
          type="submit"
          :loading="savingBacklog"
        >
          <v-icon left>mdi-chevron-left</v-icon>
          {{ $t("save") }}
          <v-icon right>mdi-chevron-right</v-icon>
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-form>
</template>
<script>
export default {
  props: ["edit"],
  watch: {
    edit: {
      handler: function () {
        if (this.edit) {
          this.backlog = this.$clone(this.edit);
        } else {
          this.backlog = this.$clone(this.emptyBacklog);
        }
      },
      immediate: true,
    },
  },
  data() {
    return {
      valid: false,
      savingBacklog: false,
      backlog: {
        Type: "user_story",
        Text: "",
        Status: "todo",
      },
      emptyBacklog: {
        Type: "user_story",
        Text: "",
        Status: "todo",
      },
      taskRules: [(v) => !!v || this.$t("please_write_something_here")],
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
  methods: {
    async saveBacklog() {
      this.$refs.form.validate();

      if (this.savingBacklog || !this.valid) {
        return false;
      }

      this.savingBacklog = true;
      if (this.edit) {
        await this.$axios.$put(
          `/api/projects/${this.$route.params.id}/backlogs/${this.edit.ID}`,
          this.backlog
        );
      } else {
        await this.$axios.$post(
          `/api/projects/${this.$route.params.id}/backlogs`,
          this.backlog
        );
      }
      this.savingBacklog = false;
      this.$emit("saved", this.$clone(this.backlog));
      this.backlog = this.$clone(this.emptyBacklog);
    },
  },
};
</script>