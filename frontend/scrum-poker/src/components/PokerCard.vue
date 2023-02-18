<template>
  <div class="card" v-bind:class="{ selected: selected }" v-on:click="select">
    <v-container>
      <v-row align="center" no-gutters>
        <v-col align="center" class="text-h5">
          <p class="noselect">{{ value  }}</p>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script setup>
const props = defineProps({
  value: {
    type: String,
    required: true
  },
  selected: {
    type: Boolean,
    required: true
  }
})
</script>

<style scoped>
  .card {
    border: 1px solid #000;
    height: 100px;
    width: 70px;
    border-radius: 5px;
  }
  .card:hover {
    background: azure;
  }
  .card:active {
    background: aquamarine;
  }
  .card.selected {
    background: aqua;
  }

  .noselect {
    -webkit-touch-callout: none; /* iOS Safari */
    -webkit-user-select: none; /* Safari */
    -khtml-user-select: none; /* Konqueror HTML */
    -moz-user-select: none; /* Firefox */
    -ms-user-select: none; /* Internet Explorer/Edge */
    user-select: none; /* Non-prefixed version, currently
                                  supported by Chrome and Opera */
  }
</style>
