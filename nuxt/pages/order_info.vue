<template lang="pug">
div
  b-row.my-1
    b-col(sm="2")
      h1 受注
    b-col(sm="10")
      p No.99999999

  b-row.my-1(align-h="center")
    b-col.text-left(sm="1")
      label(for="client_name") 顧客
    b-col(sm="4")
      b-form-select#client_name(type="text" v-model="client_name" :options="client_name_options")
    b-col.text-center(sm="1")
      label(for="employee_name") 担当者
    b-col(sm="4")
      b-form-select#employee_name(type="text" v-model="employee_name" :options="employee_name_options")

  b-row.my-1(align-h="center")
    b-col.text-left(sm="1")
      label(for="car_budget") 予算
    b-col(sm="9")
      b-form-input#car_budget(type="text" v-model="car_budget" placeholder="0")

  b-row.my-1(align-h="center")
    b-col.text-left(sm="1")
      label(for="manufacturer") メーカー
    b-col(sm="9")
      b-form-select#manufacturer(v-model="manufacturer" :options="manufacturer_options")

  b-row.my-1(align-h="center")
    b-col.text-left(sm="1")
      label(for="car_name") 車種
    b-col(sm="9")
      b-form-select#car_name(v-model="car_name" :options="car_name_options")

  b-row.my-1(align-h="center")
    b-col.text-left(sm="1")
      label(for="car_model_year") 年式
    b-col(sm="4")
      b-form-select#car_model_year(v-model="car_model_jpc_name" :options="car_model_jpc_name_options")
    b-col(sm="5")
      b-form-input#car_model_year(v-model="car_model_jpc_year")

  b-row.my-1(align-h="center")
    b-col.text-left(sm="1")
      label(for="car_model_no") 型式
    b-col(sm="9")
      b-form-select#car_model_no(v-model="car_model_no" :options="car_model_no_options")
  b-row.my-1(align-h="center")
    b-col.text-left(sm="1")
      label(for="car_exterior_color") 外装色
    b-col(sm="9")
      b-form-checkbox-group#car_exterior_color(v-model="car_exterior_color" :options="car_exterior_color_options")
  b-row.my-1(align-h="center")
    b-col.text-left(sm="1")
      label(for="car_interior_color") 内装色
    b-col(sm="9")
      b-form-checkbox-group#car_interior_color(v-model="car_interior_color" :options="car_interior_color_options")
  b-row.my-1(align-h="center")
    b-col(sm="2")
      b-button(variant="outline-success" block) 登録

</template>

<script>
export default {
  fetch({store, redirect}) {
    if (!store.state.auth) {
      return redirect('/login')
    }
  },

  data() {
    return {
      client_name  :null,
      employee_name:null,
      car_budget   :  '',
      manufacturer :  null,
      car_name     :  null,
      car_model_year: null,
      car_model_no:   null,
      car_model_jpc_name: null,
      car_model_jpc_year: null,
      car_exterior_color: [],
      car_interior_color: [],

      client_name_options: [
        { value: null, text: '顧客を選択してください' },
        { value: 'a', text: '高田１世' },
        { value: 'b', text: '高田2世' },
        { value: 'c', text: '高田3世' },
      ],
      employee_name_options: [
        { value: null, text: '担当者名を選択してください' },
        { value: 'a', text: 'ぶちお' },
        { value: 'b', text: 'ぶちこ' },
        { value: 'c', text: 'プー太郎' },
      ],
      manufacturer_options: [
        { value: null, text: 'メーカを選択してください' },
        { value: 'a', text: 'ホンダ' },
        { value: 'b', text: 'スバル' },
        { value: 'c', text: 'トヨタ' },
      ],
      car_name_options: [
        { value: null, text: '車種を選択してください' },
        { value: 'a', text: 'ホンダ' },
        { value: 'b', text: 'スバル' },
        { value: 'c', text: 'トヨタ' },
      ],
      car_model_jpc_name_options: [
        { value: null, text: '年式を選択してください' },
        { value: 'a', text: '昭和' },
        { value: 'b', text: '平成' },
      ],
      car_model_no_options:[
        { value: null, text: '型式を選択してください' },
        { value: 'a', text: '802' },
        { value: 'b', text: '320' },
        { value: 'c', text: '408' },
      ],
      car_exterior_color_options: [
        {value: '0',  text: '赤'},
        {value: '1',  text: '橙'},
        {value: '2',  text: '黄'},
        {value: '3',  text: '黄緑'},
        {value: '4',  text: '緑'},
        {value: '5',  text: '水色'},
        {value: '6',  text: '青'},
        {value: '7',  text: '紫'},
        {value: '8',  text: '白'},
        {value: '9',  text: '黒'},
        {value: '10', text: 'シルバー'},
        {value: '11', text: 'ゴールド'},
        {value: '12', text: 'その他'},
      ],
      car_interior_color_options: [
        {value: '0',  text: '赤'},
        {value: '1',  text: '橙'},
        {value: '2',  text: '黄'},
        {value: '3',  text: '黄緑'},
        {value: '4',  text: '緑'},
        {value: '5',  text: '水色'},
        {value: '6',  text: '青'},
        {value: '7',  text: '紫'},
        {value: '8',  text: '白'},
        {value: '9',  text: '黒'},
        {value: '10', text: 'シルバー'},
        {value: '11', text: 'ゴールド'},
        {value: '12', text: 'その他'},
      ]
    }
  }
}
</script>
