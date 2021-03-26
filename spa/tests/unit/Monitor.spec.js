import { createLocalVue, shallowMount } from '@vue/test-utils'
import Monitor from '@/components/Monitor.vue'
import flushPromises from 'flush-promises'

const localVue = createLocalVue()

const mockMixin = {
  methods: {
    apiGetMetrics() {
      return new Promise((resolve) => {
        resolve({
          memTotal: 300,
          memUsed: 150,
          cpuPerc: 42,
          diskTotal: 500,
          diskFree: 200,
          netBytesSent: 300,
          netBytesRecv: 300
        })
      })
    }
  }
}

describe('Monitor.vue', () => {
  it('renders monitor dials', async () => {
    const wrapper = shallowMount(Monitor, { localVue, mixins: [mockMixin] })
    await flushPromises()
    expect(wrapper).toMatchSnapshot()
  })
})
