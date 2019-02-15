<template lang="pug">
b-container
  b-row
    b-col(sm="4")
      user-nav.mt-2
    b-col(sm="8")
      b-card.mt-2(:title="$route.params.id + '&#39;s Galleries'")
      form-image(v-if="$route.params.id === $store.state.name" v-on:send="getGalleries")
      card-image(v-for="gallery in galleries" :value="gallery" :headerDisabled="true" v-on:showmodal="showModal")
      modal-image(
        :title="gallery.title"
        :src="gallery.imagepath"
        ref="modal")
</template>

<script lang="ts">
import {
  Component,
  Vue,
} from 'nuxt-property-decorator'
import axios from "axios"
import UserNav from "~/components/UserNav.vue"
import CardImage from "~/components/CardImage.vue"
import FormImage from "~/components/FormImage.vue"
import ModalImage from "~/components/ModalImage.vue"

@Component({
  components: {
    UserNav,
    CardImage,
    FormImage,
    ModalImage,
  }
})
export default class extends Vue {
  
  public galleries = [];
  public gallery = {
    title: "",
    imagepath: "",
  };
  public mounted() {
    this.getGalleries();
  }
  public getGalleries() {
    axios.get("/api/get/galleries/" + this.$route.params.id).then(({ data }) => {
      this.galleries = data.galleries;
    }).catch((result) => {
      // error
    });
  }
  public showModal(value) {
    this.gallery.title = `${value.username}(${value.datetime})`
    this.gallery.imagepath = value.imagepath;
    const modal: any = this.$refs.modal;
    modal.showModal();
  }
}
</script>
