<template lang="pug">
div
  b-row.my-1
    b-col(sm="2")
      h1 受注
    b-com(sm="10")
      p No.99999999

  b-row.my-1(align-h="center")
    b-col.text-center(sm="2")
      label(for="car_budget") 予算:
    b-col(sm="8")
      b-form-input#car_budget(type="text" v-model="car_budget" placeholder="0")
  b-row.my-1(align-h="center")
    b-col.text-center(sm="2")
      label(for="manufacturer") メーカー:
    b-col(sm="8")
      b-form-select#manufacturer(v-model="manufacturer" :options="manufacturer_options")
  b-row.my-1(align-h="center")
    b-col.text-center(sm="2")
      label(for="car_name") 車種:
    b-col(sm="8")
      b-form-select#car_name(v-model="car_name" :options="car_name_options")
  b-row.my-1(align-h="center")
    b-col.text-center(sm="2")
      label(for="car_model_year") 年式(和暦):
    b-col(sm="8")
      b-form-select#car_model_year(v-model="car_model_year" :options="car_model_year_options")

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
      car_budget   : '',
      manufacturer : null,
      car_name     : null,
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
      car_model_year_options: [
        { value: null, text: '年式を選択してください' },
        { value: 'a', text: '平成９' },
        { value: 'b', text: '平成１０' },
        { value: 'c', text: '平成１１' },
      ],

    }
  }
}
</script>
