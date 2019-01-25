<template lang="pug">
  b-container
    b-row
      b-col(sm="4")
        b-card.my-2 Categories's checkbox.
      b-col(sm="8")
        form-image
        card-image(v-for="value in values" :value="value" v-on:showmodal="showModal($event);$refs.modal.show()")
        p.text-center.mt-2(v-if="values.length === 0") Images is not uploaded yet.
    b-modal(ref="modal" size="lg" :title="value.username + '(' + value.datetime + ')'")
      b-img(
        thumnail
        fluid
        :src="value.imagepath")
</template>

<script lang="ts">
import {
  Component,
  Vue
} from "nuxt-property-decorator"
import FormImage from "~/components/FormImage.vue"
import CardImage from "~/components/CardImage.vue"
import axios from "axios"


@Component({
  components: {
    FormImage,
    CardImage
  }
})
export default class extends Vue {
  public values = [];
  public value = {
    username:  "",
    evalparam: null,
    evalsum:   null,
    favorite:  null,
    comment:   null,
    number:    0,
    imagepath: "",
    datetime:  "",
  };

  public mounted() {
    axios.get("/api/get/galleries").then(({ data }) => {
      if (data.galleries) {
        this.values = data.galleries;
      }
    });
  }

  public showModal(value) {
    this.value = value;
  }
}
</script>

<style lang="sass" scoped></style>
