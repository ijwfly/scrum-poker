<template>
  <v-table>
    <thead>
      <tr>
        <th class="text-left">Name</th>
        <th class="text-left">Estimate</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="(estimate, index) in estimates" :key="index">
        <td>{{ estimate.name }}</td>
        <td v-if="revealed">{{ estimate.estimate }}</td>
        <td v-else-if="estimate.estimate !== null">X</td>
        <td v-else>-</td>
      </tr>
    </tbody>
  </v-table>
</template>

<script>
export default {
  name: "ResultsTable",
  props: {
    estimates: {
      type: Array,
      required: true,
    },
    revealed: {
      type: Boolean,
      required: true,
    }
  },
}
</script>

<style scoped>

</style>
