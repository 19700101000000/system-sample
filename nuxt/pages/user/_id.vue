<template lang="pug">
b-container
  b-row
    b-col(sm="4")
      user-nav.mt-2
    b-col(sm="8")
      b-card.mt-2(
        v-if="$store.state.name === $route.params.id"
        title="My Page") Please select from left nav.
      wanted-list-item.mt-2(
        v-else
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
          :state="stateTitle"
          :disabled="disabled")
        p.text-right {{ title.length }} / 100
      b-form-group(
        label="Description:"
        label-for="inputDescription")
        b-form-textarea#inputDescription(
          placeholder="Please input Description"
          v-model="description"
          :rows="6"
          :state="stateDescription"
          :disabled="disabled")
        p.text-right {{ description.length}} / 500
      b-form-group(
        label="Price(JPY):"
        label-for="inputPrice")
        b-form-input#inputPrice(
          :class="[price.length > 0 ? 'text-right' : '']"
          v-model="price"
          :state="statePrice"
          :formatter="priceFormat"
          placeholder="Please input price."
          :disabled="disabled")
        p.text-right more than JPY {{ price2Text(value.price) }}
    b-card(
      title="WANTED INFO")
      b-row
        b-col(sm="2") Title:
        b-col(sm="10") {{ value.title }}
      b-row
        b-col(sm="2") Description:
        b-col(sm="10") {{ value.description }}
      b-row
        b-col(sm="2") Price:
        b-col(sm="10") {{ price2Text(value.price) }}
    p.text-danger(v-if="error") {{ errorMsg }}
    div(slot="modal-footer")
      b-button(
        variant="danger"
        v-on:click="onClickCheat"
      ) プレゼン用チートボタン
      b-button.ml-2(
        variant="outline-secondary"
        v-on:click="modalCancel"
        :disabled="disabled") Cancel
      b-button.ml-2(
        variant="outline-primary"
        :disabled="!stateTitle || !stateDescription || !statePrice || disabled"
        v-on:click="modalOK") Send
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
  public title       = "";
  public description = "";
  public price       = "";
  public priceTrue   = 0;
  public disabled    = false;

  public wanteds = [];
  public value = {
    username: "",
    number: 0,
    title: "",
    description: "",
    price: 0,
  }

  public error = false;
  public errorMsg = "Fail uploaded.";

  public get stateTitle(): boolean {
    return this.title.length > 0 && this.title.length <= 100;
  }
  public get stateDescription(): boolean {
    return this.description.length > 0 && this.title.length <= 500;
  }
  public get statePrice(): boolean {
    return this.price.length > 0 && this.priceTrue >= this.value.price;
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
  public price2Text(price: number): string {
    return price.toString().replace(/(\d)(?=(\d\d\d)+(?!\d))/g, "$1,");
  }

  public onClickCheat(event: Event) {
    this.title = "Twitter用アイコンを描いてほしいです";
    this.description = "アイコン画像を描いてほしいです\nお願いします";
    this.price = "1000000";
  }

  public modalCancel(event: Event) {
    let modal: any = this.$refs.modal;
    modal.hide();
  }
  public modalOK(event: Event) {
    this.disabled = true;
    axios.post("/api/upload/request", {
      ownername: this.value.username,
      wanted: this.value.number,
      title: this.title,
      description: this.description,
      price: this.priceTrue,
    }).then((result) => {
      this.title = "";
      this.description = "";
      this.price = "";
      this.error = false;
      this.disabled = false;
      let modal: any = this.$refs.modal;
      modal.hide();
    }).catch((result) => {
      this.error = true;
      this.disabled = false;
    });
  }
  public mounted() {
    axios.get("/api/get/works/wanteds/" + this.$route.params.id).then(({ data }) => {
      this.wanteds = data.wanteds;
    });
  }
  public showmodal(value) {
    this.value = value;
    this.price = this.value.price.toString();
    let modal: any = this.$refs.modal;
    modal.show();
  }
}
</script>

<style lang="sass"></style>
