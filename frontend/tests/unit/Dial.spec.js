import { mount } from '@vue/test-utils'
import Dial from '@/components/Dial.vue'

describe('Dial.vue', () => {
  it('renders a dial', async () => {
    const wrapper = mount(Dial, {
      propsData: {
        value: 22,
        title: 'Test Dial'
      }
    })

    expect(wrapper.html()).toMatchSnapshot()
  })
})
