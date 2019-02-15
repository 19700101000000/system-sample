<template lang="pug">
  b-modal(
    ref="modal"
    size="lg"
    ok-only
    ok-variant="outline-primary"
    :title="title")
    div.text-center
      b-img(
        fluid
        :src="src")
</template>
<script lang="ts">
import {
  Vue,
  Prop,
  Component,
} from "nuxt-property-decorator"

@Component({})
export default class ModalImage extends Vue {
  @Prop() title!: string;
  @Prop() src!: string;

  public showModal(): void {
    const modal: any = this.$refs.modal;
    modal.show();
  }
}
</script>
