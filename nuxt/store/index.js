import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'
import Cookie from 'js-cookie'

Vue.use(Vuex)

var cookieparser = require('cookieparser')

const store = () => new Vuex.Store({

  state: {
    auth: {
      user: null
    }
  },

  mutations: {
    SET_USER: function (state, user) {
      state.auth.user = user
    }
  },

  actions: {
    nuxtServerInit ({ commit }, { req }) {
      let auth = null
      if (req.headers.cookie) {
        var parsed = cookieparser.parse(req.headers.cookie)
        auth = JSON.parse(parsed.auth || {
          user: null
        })
      }

      commit('SET_USER', auth.user)
    },
    // TODO create login and logout methods
    async login({ commit }, { username, password }) {
      // const params = new URLSearchParams()
      // params.append('username', username)
      // params.append('password', password)

      try {
        const { data } = await axios.post('/api/login', {
          username: username,
          password: password
        })
        console.log(data)
      } catch(error) {
        console.log(error.message)
      }
      commit('SET_USER', username)
      const auth = {
        user: username
      }
      Cookie.set('auth', auth)
    },
    async logout({ commit }) {
      try {
        const { data } = await axios.post('/api/logout')
        console.log(data)
      } catch(error) {
        console.log(error.message)
      }
      commit('SET_USER', null)
      const auth = {
        user: null
      }
      Cookie.set('auth', auth)
    }
  }
})

export default store

