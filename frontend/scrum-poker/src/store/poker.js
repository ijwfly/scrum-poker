import { defineStore } from 'pinia'

export const usePokerStore = defineStore('poker', {
  state: () => {
    return {
      estimateOptions: ['?', 0.5, 1, 2, 3, 5, 8, 13],
      selectedEstimateIndex: null,
      selectedEstimateValue: null,

      estimatesRevealed: false,
      estimates: [
        { name: 'John', estimate: 1 },
        { name: 'Jenny', estimate: 13 },
        { name: 'Jane', estimate: null },
        { name: 'Jenny', estimate: 8 },
        { name: 'Jack', estimate: 3 },
        { name: 'Jill', estimate: null },
        { name: 'Maxim', estimate: null},
      ]
    }
  },
  actions: {
    selectEstimate(index) {
      this.selectedEstimateIndex = index
      this.selectedEstimateValue = this.estimateOptions[index]

      this.estimates[6].estimate = this.estimateOptions[index]
    },
    deleteEstimates() {
      this.selectedEstimateIndex = null
      this.selectedEstimateValue = null
      this.estimatesRevealed = false
      for (let i = 0; i < this.estimates.length; i++) {
        this.estimates[i].estimate = null
      }
    },
    revealEstimates() {
      this.estimatesRevealed = !this.estimatesRevealed
    }
  }
})
