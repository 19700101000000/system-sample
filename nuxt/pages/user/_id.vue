<template lang="pug">
b-container
  b-row
    b-col(sm="4")
      user-nav.mt-2
    b-col(sm="8")
      wanted-list-item.mt-2(
        v-for="wanted in wanteds"
        :value="wanted"
        v-on:showmodal="showmodal")
  b-modal(ref="modal" size="lg" title="New Request")
    b-form
      b-form-group(
        label="Title:"
        label-for="inputTitle")
        b-form-input#inputTitle(
          placeholder="Please input Title"
          v-model="title"
          :state="stateTitle")
      b-form-group(
        label="Description:"
        label-for="inputDescription")
        b-form-textarea#inputDescription(
          placeholder="Please input Description"
          v-model="description"
          :rows="6"
          :state="stateDescription")
      b-form-group(
        label="Price(JPY):"
        label-for="inputPrice")
        b-form-input#inputPrice(
          :class="[price.length > 0 ? 'text-right' : '']"
          v-model="price"
          :state="statePrice"
          :formatter="priceFormat"
          placeholder="Please input price.")
    p {{ value.username }}: {{ value.title }}
    p {{ value.description }}
    p JPY {{ value.price }}
    div(slot="modal-footer")
      b-button(variant="outline-secondary") Cancel
      b-button.ml-2(variant="outline-primary") Send
</template>

<script lang="ts">
import {
  Component,
  Vue
} from 'nuxt-property-decorator'
import axios from "axios"
import UserNav from "~/components/UserNav.vue"
import WantedListItem from "~/components/WantedListItem.vue"

@Component({
  components: {
    UserNav,
    WantedListItem
  }
})
export default class extends Vue {
  public title = "";
  public description = "";
  public price       = "";
  public priceTrue   = 0;

  public get stateTitle(): boolean {
    return this.title.length > 0;
  }
  public get stateDescription(): boolean {
    return this.description.length > 0;
  }
  public get statePrice(): boolean {
    return this.price.length > 0;
  }
  public priceFormat(value: string, event): string {
    const match = value.replace(/,/g, '').match(/\d+/);
    if (match) {
      const price = match[0].replace(/(^0+)(\d+)/, '$2');
      this.priceTrue = parseInt(price);

      return price.replace(/(\d)(?=(\d\d\d)+(?!\d))/g, "$1,");
    }
    return "";
  }

  public wanteds = [];
  public value = {
    username: "",
    number: 0,
    title: "",
    description: "",
    price: 0,
  }
  public mounted() {
    axios.get("/api/get/works/wanteds/" + this.$route.params.id).then(({ data }) => {
      this.wanteds = data.wanteds;
    });
  }
  public showmodal(value) {
    this.value = value;
    let modal: any = this.$refs.modal;
    modal.show();
  }
}
</script>

<style lang="sass"></style>
