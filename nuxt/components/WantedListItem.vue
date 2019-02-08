<template lang="pug">
  b-card(
    :border-variant="cardVariant"
    no-body)
    b-card-body
      h4 {{ value.title }}
      p {{ value.description }}
      p.card-text JPY {{ value.price }}
    div(slot="footer")
      div(v-if="$store.state.name === value.username")
        template(v-if="value.allrequests") All requests: {{ value.allrequests }}
        template(v-else) No Requests.
      div.text-right
        b-button(
          v-if="$store.state.name !== value.username"
          :disabled="$store.state.name === ''"
          variant="outline-primary"
          size="sm"
          v-on:click="$emit('showmodal', value)") new request
        b-button(
          v-else
          variant="outline-secondary"
          size="sm"
          v-on:click="$emit('showmodalEdit', value)") edit
        b-button.ml-2(
          v-if="$store.state.name === value.username"
          :variant="requestsVariant"
          size="sm"
          v-on:click="$emit('showmodalRequests', value)") requests
          template(v-if="value.requests > 0") ({{ value.requests }})
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
  alive:       boolean,
  requests:    number,
  allrequests: number,
}

@Component({})
export default class WantedListItem extends Vue {
  @Prop() value: WantedValue;

  public get cardVariant(): string {
    if (this.value.alive) {
      return "success";
    }
    return "danger";
  }
  public get requestsVariant(): string {
    if (this.value.requests > 0) {
      return "outline-success";
    }
    return "outline-primary";
  }
}
</script>
