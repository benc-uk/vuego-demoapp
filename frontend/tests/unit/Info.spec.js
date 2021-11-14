import { mount, flushPromises } from '@vue/test-utils'
import Info from '@/components/Info.vue'

global.fetch = jest.fn(() =>
  Promise.resolve({
    ok: true,
    status: 200,
    statusText: 'OK',
    headers: {
      get: () => 'application/json'
    },
    json: () =>
      Promise.resolve({
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
)

describe('Info.vue', () => {
  it('renders info screen', async () => {
    const wrapper = mount(Info, {})
    await flushPromises()
    expect(wrapper.html()).toMatchSnapshot()
  })
})
