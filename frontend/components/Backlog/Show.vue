<template>
  <v-card>
    <v-card-title class="mt-4 mb-2 flex-column">
      <div class="text-headline primary--text mr-2">
        {{ $backlogNumber(backlog) }}
      </div>
      <div class="text-body-1 text-center" style="word-break: normal">
        {{ $backlogEmoji(backlog.Type) }} {{ backlog.Text }}
      </div>
      <div class="text-center">
        <v-btn outlined rounded text small @click="editDialog = true">{{
          $t("edit")
        }}</v-btn>
        <v-dialog
          v-model="editDialog"
          max-width="800px"
          style="overflow: hidden"
        >
          <BacklogSave @saved="onSave" :edit="backlog"></BacklogSave>
        </v-dialog>
      </div>
      <div class="mt-4" style="width: 50%">
        <v-divider></v-divider>
      </div>
    </v-card-title>
    <v-card-text>
      <v-simple-table>
        <template v-slot:default>
          <thead>
            <tr>
              <th>{{ $t("task") }}</th>
              <th>{{ $t("point") }}</th>
              <th style="max-width: 30px">{{ $t("status") }}</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="subtask in subtasks" :key="subtask.ID">
              <td>{{ subtask.Text }}</td>
              <td>
                {{ subtask.Point }}
              </td>
              <td style="max-width: 30px" class="p-0">
                <Status :status="subtask.Status" type="subtask"></Status>
              </td>
            </tr>
          </tbody>
        </template>
      </v-simple-table>
    </v-card-text>
  </v-card>
</template>
<script>
export default {
  props: ["backlog"],
  data() {
    return {
      editDialog: false,
      loadingSubtasks: false,
      subtasks: [],
    };
  },
  watch: {
    backlog: {
      handler: function () {
        this.loadSubtasks();
      },
      immediate: true,
    },
  },
  methods: {
    onSave(backlog) {
      this.$emit("update:backlog", backlog);
      console.log(backlog)
      this.editDialog = false;
    },
    async loadSubtasks() {
      this.loadingSubtasks = true;
      this.subtasks = await this.$axios.$get(
        `/api/projects/${this.$route.params.id}/backlogs/${this.backlog.ID}/subtasks`
      );
      this.loadingSubtasks = false;
    },
  },
};
</script>