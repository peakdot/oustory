 <template>
  <div>
    <template v-if="!loadingBacklogs && backlogs.length > 0">
      <v-simple-table>
        <template v-slot:default>
          <thead>
            <tr>
              <th class="text-left">{{ $t("#") }}</th>
              <th class="text-left">{{ $t("task") }}</th>
              <th class="text-left">{{ $t("status") }}</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="backlog in backlogs" :key="backlog.ID">
              <td
                class="primary--text font-weight-medium"
                @click="$emit('select', backlog)"
                style="cursor: pointer"
              >
                {{ $backlogNumber(backlog) }}
              </td>
              <td @click="$emit('select', backlog)" style="cursor: pointer">
                {{ $backlogEmoji(backlog.Type) }}
                {{ backlog.Text }}
              </td>
              <td>
                <Status :status="backlog.Status" type="backlog"></Status>
              </td>
            </tr>
          </tbody>
        </template>
      </v-simple-table>
    </template>
    <template v-if="!loadingBacklogs && backlogs.length == 0">
      <v-card-text class="text-center">
        <v-img src="/no_result.png" max-height="64" contain />
        No variance energy.
      </v-card-text>
    </template>
    <template v-if="loadingBacklogs">
      <div class="text-center">
        <v-progress-circular
          indeterminate
          color="primary"
          class="my-4"
        ></v-progress-circular>
      </div>
    </template>
  </div>
</template>
<script>
export default {
  data() {
    return {
      backlogs: [],
      loadingBacklogs: false,
    };
  },
  mounted() {
    this.loadBacklogs();
  },
  methods: {
    async loadBacklogs() {
      this.loadingBacklogs = true;
      this.backlogs = await this.$axios.$get(
        `/api/projects/${this.$route.params.id}/backlogs`
      );
      this.loadingBacklogs = false;
    },
  },
};
</script>
