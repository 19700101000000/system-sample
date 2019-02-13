<template lang="pug">
b-container
  b-row
    b-col(sm="4")
      user-nav.mt-2
    b-col(sm="8" v-if="$store.state.name !== ''")
      b-card.mt-2(no-body)
        b-card-body
          h4 MY WANTEDS
          p.card-text.text-right
            b-button(variant="outline-primary" v-b-modal.newWanted) new wanted
      wanted-list-item.mt-2(
        v-for="wanted in wanteds"
        :value="wanted"
        v-on:showmodalRequests="showmodalRequests"
        v-on:showmodalEdit="showmodalEdit")

  b-modal(
    title="Requests"
    ref="modalRequests"
    size="lg")
    p(v-if="requests.length === 0") No Requests.
    request-list-item.mt-2(v-for="request in requests" :value="request" v-on:checked="onCheckedRequest")

  b-modal(
    title="Edit"
    ref="modalEdit"
    v-on:ok="editOK"
    v-on:hide="editHide"
    ok-variant="outline-primary"
    ok-title="Apply"
    cancel-variant="outline-secondary"
    cancel-title="Cancel"
    size="lg")
    b-card(
      :border-variant="editBorderVariant"
      no-body)
      b-card-body
        h4 {{ wantedValue.title }}
        p.card-text {{ wantedValue.description }}
        p.card-text JPY {{ wantedValue.price }}
    b-form
      b-form-group(label="Active:")
        b-form-radio-group(
          buttons
          :button-variant="[wantedValue.alive? 'outline-success' : 'outline-danger']"
          v-model="wantedValue.alive" name="active")
          b-form-radio(:value="true") Yes
          b-form-radio(:value="false") No

  b-modal#newWanted(
    title="New Wanted"
    size="lg"
    ref="newWanted")
    b-form
      b-form-group(
        label="Title:"
        label-for="inputTitle")
        b-form-input#inputTitle(
          v-model="title"
          :state="stateTitle"
          placeholder="Please input title."
          :disabled="modalDisabled")
        p.text-right {{ title.length }} / 100
      b-form-group(
        label="Description:"
        label-for=inputDescription)
        b-form-textarea#inputDescription(
          v-model="description"
          :state="stateDescription"
          placeholder="Please input description."
          :rows="6"
          :disabled="modalDisabled")
        p.text-right {{ description.length }} / 500
      b-form-group(
        label="Price(JPY):"
        label-for="inputPrice")
        b-form-input#inputPrice(
          :class="[price.length > 0 ? 'text-right' : '']"
          v-model="price"
          :state="statePrice"
          :formatter="priceFormat"
          placeholder="Please input price."
          :disabled="modalDisabled")
    div.text-danger(v-if="error") {{ errorMsg }}
    div(slot="modal-footer")
      b-button(
        slot="modal-cancel"
        variant="outline-secondary"
        v-on:click="$refs.newWanted.hide()") Cancel
      b-button.ml-2(
        slot="modal-ok"
        variant="outline-primary"
        v-on:click="modalOK"
        :disabled="!stateTitle || !stateDescription || !statePrice || modalDisabled") Send
</template>

<script lang="ts">
import {
  Component,
  Vue
} from "nuxt-property-decorator"
import axios from "axios"
import UserNav from "~/components/UserNav.vue"
import WantedListItem from "~/components/WantedListItem.vue"
import RequestListItem from "~/components/RequestListItem.vue"
import requestVue from "./request.vue";


@Component({
  components: {
    UserNav,
    WantedListItem,
    RequestListItem,
  }
})
export default class extends Vue {
  public modalDisabled = false;

  public error = false;
  public errorMsg = "Fail Uploaded.";
  
  public title       = "";
  public price       = "10,000";
  public priceTrue   = 0;
  public description = "";

  public wanteds = [];
  public requests = [];
  public wantedValue = {
    username:    "",
    number:      0,
    title:       "",
    description: "",
    price:       0,
    alive:       true,
    requests:    0,
  };
  public edit = {
    alive: false,
  };

  public priceFormat(value: string, event): string {
    const match = value.replace(/,/g, '').match(/\d+/);
    if (match) {
      const price = match[0].replace(/(^0+)(\d+)/, '$2');
      this.priceTrue = parseInt(price);

      return price.replace( /(\d)(?=(\d\d\d)+(?!\d))/g, "$1,");
    }
    return "";
  }

  public get editBorderVariant(): string {
    if (this.wantedValue.alive) {
      return "success";
    }
    return "danger";
  }
  public get stateTitle(): boolean {
    return this.title.length > 0 && this.title.length <= 100;
  }
  public get stateDescription(): boolean {
    return this.description.length > 0 && this.description.length <= 500;
  }
  public get statePrice(): boolean {
    return this.price.length > 0;
  }

  public editOK() {
    this.edit.alive = this.wantedValue.alive;
  }
  public editHide() {
    this.wantedValue.alive = this.edit.alive;
  }
  public modalOK() {
    this.modalDisabled = true;
    axios.post("/api/upload/wanted", {
      title: this.title,
      description: this.description,
      price: this.priceTrue,
    }).then((result) => {
      this.modalDisabled = false;
      this.title = "";
      this.description = "";
      this.price = "10000";
      this.priceTrue = 10000;
      this.error = false;
      let modal: any = this.$refs.newWanted;
      modal.hide();
    }).catch((result) => {
      this.error = true;
      this.modalDisabled = false;
    });
  }

  public onCheckedRequest() {
    this.wantedValue.requests -= 1;
  }

  public showmodalRequests(value) {
    this.wantedValue = value;
    axios.get("/api/get/works/requests/" + this.wantedValue.number).then(({ data }) => {
      this.requests = data.requests;
    }).catch((result) => {
      this.requests = [];
    });
    let modal: any = this.$refs.modalRequests;
    modal.show();
  }
  public showmodalEdit(value) {
    this.wantedValue = value;
    this.edit.alive = this.wantedValue.alive;
    let modal: any = this.$refs.modalEdit;
    modal.show();
  }
  public mounted() {
    axios.get("/api/get/works/wanteds").then(({ data }) => {
      this.wanteds = data.wanteds;
    });
  }
}
</script>

<style lang="sass"></style>
