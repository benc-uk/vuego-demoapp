import { mount, flushPromises } from '@vue/test-utils'
import Monitor from '@/components/Monitor.vue'

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
        memTotal: 300,
        memUsed: 150,
        cpuPerc: 42,
        diskTotal: 500,
        diskFree: 200,
        netBytesSent: 300,
        netBytesRecv: 300
      })
  })
)

describe('Monitor.vue', () => {
  it('renders monitor screen', async () => {
    const wrapper = mount(Monitor, {})
    await flushPromises()
    expect(wrapper.html()).toMatchSnapshot()
  })
})
