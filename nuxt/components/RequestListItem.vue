<template lang="pug">
  b-card(:title="value.title") {{ value.description }}
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
}
</script>
