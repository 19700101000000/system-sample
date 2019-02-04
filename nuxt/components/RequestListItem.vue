<template lang="pug">
  b-card(no-body :border-variant="borderVariant")
    div(slot="header")
      template(v-if="!value.establish && value.alive") 承認待ち
      template(v-else-if="value.establish && value.alive") 成立
      template(v-else-if="value.establish && !value.alive") 終了
      template(v-else) 破談
    b-card-body
      h4 {{ value.title }}
      p.card-text {{ value.description }}
    div(slot="footer") JPY {{ value.price }}
</template>

<script lang="ts">
import {
  Component,
  Prop,
  Vue
} from "nuxt-property-decorator"

interface RequestValue {
  ownername:   string,
  wanted:      number,
  number:      number,
  username:    string,
  title:       string,
  description: string,
  price:       number,
  establish:   boolean,
  alive:       boolean,
}
@Component
export default class RequestListItem extends Vue {
  @Prop() value: RequestValue;
  public get borderVariant():string {
    if (!this.value.establish && this.value.alive) {
      return "primary";
    } else if (this.value.establish && this.value.alive) {
      return "success";
    } else if (this.value.establish && !this.value.alive) {
      return "secondary"
    }
    return "danger";
  }
}
</script>
