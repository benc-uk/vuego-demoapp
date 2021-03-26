import { createLocalVue, shallowMount } from '@vue/test-utils'
import Info from '@/components/Info.vue'
import flushPromises from 'flush-promises'

const localVue = createLocalVue()

const mockMixin = {
  methods: {
    apiGetInfo() {
      return new Promise((resolve) => {
        resolve({
          hostname: 'test',
          isContainer: true,
          isKubernetes: true,
          platform: 'PDP-11',
          architecture: '98 bit',
          os: 'MegaOS: 3000',
          mem: 1234567890,
          envVars: ['UNIT_TESTS=Are pointless']
        })
      })
    }
  }
}

describe('Info.vue', () => {
  it('renders info screen', async () => {
    const wrapper = shallowMount(Info, { localVue, mixins: [mockMixin] })
    await flushPromises()
    expect(wrapper).toMatchSnapshot()
  })
})
