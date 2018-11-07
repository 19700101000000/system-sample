import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'
import Cookie from 'js-cookie'

Vue.use(Vuex)

var cookieparser = require('cookieparser')

const store = () => new Vuex.Store({

  state: {
    auth: null
  },

  mutations: {
    update: function (state, auth) {
      state.auth = auth
    }
  },

  actions: {
    nuxtServerInit ({ commit }, { req }) {
      let auth = null
      if (req.headers.cookie) {
        var parsed = cookieparser.parse(req.headers.cookie)
        if (!parsed.auth) {
          return
        }
        auth = JSON.parse(parsed.auth)
      }

      commit('update', auth)
    },

    async login({ commit }, { username, password }) {
      try {
        const { data } = await axios.post('/api/auth', {
          username: username,
          password: password
        })
        if(!data.success) {
          return
        }
        commit('update', data.auth)
        Cookie.set('auth', data.auth)

      } catch(error) {
        console.log(error.message)
        return
      }
      window.location.href = '/'
    },

    async logout({ commit }) {
      commit('update', null)
      Cookie.set('auth', null)
      window.location.href = '/'
    }
  }
})

export default store

