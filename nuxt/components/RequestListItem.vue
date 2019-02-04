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
      b-card(
        v-if="wanted"
        title="wanted")
        b-row
          b-col.text-secondary(sm="3") Owner:
          b-col(sm="9") {{ value.wanted.username }}
        b-row
          b-col.text-secondary(sm="3") Title:
          b-col(sm="9") {{ value.wanted.title }}
        b-row
          b-col.text-secondary(sm="3") Description:
          b-col(sm="9") {{ value.wanted.description }}
        b-row
          b-col.text-secondary(sm="3") Price:
          b-col(sm="9") JPY {{ value.wanted.price }}
    div(slot="footer") JPY {{ value.price }}
      div.text-right
        b-button.ml-2(variant="outline-secondary") 保留
        b-button.ml-2(variant="outline-primary") 交渉
        b-button.ml-2(variant="outline-success") 承諾
        b-button.ml-2(variant="outline-danger") 破棄
</template>

<script lang="ts">
import {
  Component,
  Prop,
  Vue
} from "nuxt-property-decorator"

interface RequestValue {
  wanted: {
    username:    string,
    number:      number,
    title:       string,
    description: string,
    price:       number,
    alive:       boolean,
  },
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
  @Prop() wanted: boolean;

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
