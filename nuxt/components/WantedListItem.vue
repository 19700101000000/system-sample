<template lang="pug">
  b-card(no-body)
    b-card-body
      h4 {{ value.title }}
      p {{ value.description }}
      p.card-text JPY {{ value.price }}
        b-button.ml-2(variant="outline-primary" size="sm") new request
</template>

<script lang="ts">
import {
  Component,
  Prop,
  Vue
} from "nuxt-property-decorator"

interface WantedValue {
  username:    string,
  number:      number,
  title:       string,
  description: string,
  price:       number,
}

@Component({})
export default class WantedListItem extends Vue {
  @Prop() value: WantedValue;
}
</script>
