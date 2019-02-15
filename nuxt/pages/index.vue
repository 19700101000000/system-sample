<template lang="pug">
  b-container
    b-row
      b-col(sm="4")
        b-card.my-2 Categories's checkbox.
      b-col(sm="8")
        form-image(v-on:send="getGalleries")
        card-image(v-for="value in values" :value="value" v-on:showmodal="showModal")
        p.text-center.mt-2(v-if="values.length === 0") Images is not uploaded yet.
        modal-image(
          :title="gallery.title"
          :src="gallery.imagepath"
          ref="modal")
</template>

<script lang="ts">
import {
  Component,
  Vue
} from "nuxt-property-decorator"
import FormImage from "~/components/FormImage.vue"
import CardImage from "~/components/CardImage.vue"
import axios from "axios"
import ModalImage from "~/components/ModalImage.vue"

@Component({
  components: {
    FormImage,
    CardImage,
    ModalImage,
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
  public gallery = {
    title: "",
    imagepath: "",
  }

  public mounted() {
    this.getGalleries();
  }
  public getGalleries() {
    axios.get("/api/get/galleries").then(({ data }) => {
      if (data.galleries) {
        this.values = data.galleries;
      }
    });
  }

  public showModal(value) {
    this.value = value;
    this.gallery.imagepath = this.value.imagepath;
    this.gallery.title = `${this.value.username}(${this.value.datetime})`;

    const modal: any = this.$refs.modal;
    modal.showModal();
  }
}
</script>

<style lang="sass" scoped></style>
