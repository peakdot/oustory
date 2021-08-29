<template>
  <div>
    <v-row>
      <v-col>
        <v-btn rounded color="primary" @click="addDialog = true">
          <v-icon left>mdi-plus</v-icon>
          {{ $t("add_backlog") }}
        </v-btn>
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <v-card class="mx-auto rounded-xl" style="overflow: hidden">
          <BacklogList @select="onSelect"></BacklogList>
        </v-card>
      </v-col>
    </v-row>

    <v-dialog v-model="backlogDialog" max-width="800px">
      <BacklogShow :backlog.sync="selectedBacklog"></BacklogShow>
    </v-dialog>
    <v-dialog v-model="addDialog" max-width="800px" style="overflow: hidden">
      <BacklogSave @saved="onSave"></BacklogSave>
    </v-dialog>
  </div>
</template>
<script>
export default {
  data() {
    return {
      backlogDialog: false,
      addDialog: false,
      selectedBacklog: null,
    };
  },
  methods: {
    onSelect(backlog) {
      this.backlogDialog = true;
      this.selectedBacklog = backlog;
    },
    onSave() {
      this.addDialog = false;
    },
  },
};
</script>