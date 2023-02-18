<template>
  <v-container>
    <v-container>
      <v-row>
        <v-col v-for="(value, index) in poker.estimateOptions" :key="index">
          <PokerCard :value="value" :selected="poker.selectedEstimateIndex === index" v-on:click="poker.selectEstimate(index)"></PokerCard>
        </v-col>
      </v-row>
    </v-container>
    <v-divider></v-divider>
    <v-container>
      <v-row>
        <v-col align="center">
          <v-btn @click="poker.deleteEstimates">Delete estimates</v-btn>
        </v-col>
        <v-col align="center">
          <v-btn @click="poker.revealEstimates">Reveal estimates</v-btn>
        </v-col>
      </v-row>
    </v-container>
  </v-container>
  <v-container>
    <ResultsTable :estimates="poker.estimates" :revealed="poker.estimatesRevealed"></ResultsTable>
  </v-container>
</template>

<script setup>
import PokerCard from "@/components/PokerCard";
import {usePokerStore} from "@/store/poker";
import ResultsTable from "@/components/ResultsTable";

const poker = usePokerStore();

</script>

<style scoped>

</style>
