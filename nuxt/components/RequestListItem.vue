<template lang="pug">
b-card(no-body :border-variant="borderVariant")
  div(slot="header")
    template(v-if="!value.establish && value.alive") 承認待ち
    template(v-else-if="value.establish && value.alive") 成立
    template(v-else-if="value.establish && !value.alive") 終了
    template(v-else) 破談
  b-card-body
    b-row
      b-col(sm="3") Requester:
      b-col(sm="9")
        b-link(:href="`/user/${value.username}`") {{ value.username }}
    b-row
      b-col(sm="3") Title:
      b-col(sm="9") {{ value.title }}
    b-row
      b-col(sm="3") Description:
      b-col(sm="9") {{ value.description }}
    b-row
      b-col(sm="3") Price:
      b-col(sm="9") {{ getPrice }}
    b-card(
      v-if="wanted"
      title="wanted")
      b-row
        b-col.text-secondary(sm="3") Owner:
        b-col(sm="9")
          b-link(:href="`/user/${value.wanted.username}`") {{ value.wanted.username }}
      b-row
        b-col.text-secondary(sm="3") Title:
        b-col(sm="9") {{ value.wanted.title }}
      b-row
        b-col.text-secondary(sm="3") Description:
        b-col(sm="9") {{ value.wanted.description }}
      b-row
        b-col.text-secondary(sm="3") Price:
        b-col(sm="9") JPY {{ getWantedPrice }}
  div.text-right(v-if="$store.state.name === value.wanted.username && value.alive" slot="footer")
    b-button.ml-2(variant="outline-secondary" v-if="!value.check" v-on:click="onClickChecked") 保留
    b-button.ml-2(variant="outline-primary" disabled) 交渉
    template(v-if="value.establish")
      b-button.ml-2(variant="outline-secondary" v-on:click="onClickFinish") 終了
    template(v-else)
      b-button.ml-2(variant="outline-success" v-on:click="onClickSuccess") 承諾
      b-button.ml-2(variant="outline-danger" v-on:click="onClickDanger") 破棄
</template>

<script lang="ts">
import {
  Component,
  Prop,
  Vue
} from "nuxt-property-decorator"
import axios from "axios"

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
  check:       boolean,
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

  public get getPrice(): string {
    return this.value.price.toString().replace(/(\d)(?=(\d\d\d)+(?!\d))/g, "$1,");
  }
  public get getWantedPrice(): string {
    return this.value.wanted.price.toString().replace(/(\d)(?=(\d\d\d)+(?!\d))/g, "$1,");
  }

  public onClickChecked(event: Event) {
    this.value.check = true;
    axios.post("/api/update/request/checked", {
      wanted: this.value.wanted.number,
      request: this.value.number,
    }).then((result) => {
      this.$emit("checked");
    }).catch((result) => {
      this.value.check = false;
    })
  }
  public onClickSuccess(event: Event) {
    this.postStatus(true, true);
  }
  public onClickDanger(event: Event) {
    this.postStatus(false, false);
  }
  public onClickFinish(event: Event) {
    this.postStatus(true, false);
  }
  private postStatus(establish: boolean, alive: boolean) {
    let establishBack = this.value.establish;
    let aliveBack     = this.value.alive;
    this.value.establish = establish;
    this.value.alive = alive;
    axios.post("/api/update/request/status", {
      wanted:    this.value.wanted.number,
      request:   this.value.number,
      establish: this.value.establish,
      alive:     this.value.alive,
    }).then((result) => {
      if(!this.value.check) {
        this.$emit("checked");
      }
    }).catch((result) => {
      this.value.establish = establishBack;
      this.value.alive = aliveBack;
    });
  }
}
</script>
