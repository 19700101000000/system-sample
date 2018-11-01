<template lang="pug">
b-table(striped hover :items="items")
</template>

<script lang="ts">
import {
  Component,
  Vue,
} from "nuxt-property-decorator"

@Component({})
export default class RankTable extends Vue {
  items: object[] = [
    { rank: 1 },
    { rank: 2 },
    { rank: 3 },
    { rank: 4 },
    { rank: 5 },
  ];
}
</script>
