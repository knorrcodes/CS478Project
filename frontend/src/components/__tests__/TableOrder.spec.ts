import Vue from "vue";
import { localVue } from "@t/setup-test";
import { mount, shallowMount } from "@vue/test-utils";
import { expect } from "chai";

import { genTestOrder } from "@t/api-data";
import TableOrder from "../TableOrder.vue";

describe("TableOrder", () => {
  it("correctly displays the customer code", () => {
    const wrapper = shallowMount(TableOrder, {
      localVue,
      propsData: {
        currentOrder: genTestOrder(true),
        refetchFunc() {}
      }
    });

    expect(wrapper.text()).contains("Code: TEST12");
  });

  it("correctly list the order items", () => {
    const wrapper = shallowMount(TableOrder, {
      localVue,
      propsData: {
        currentOrder: genTestOrder(),
        refetchFunc() {}
      }
    });

    expect(wrapper.text()).contains("Test Product 1 $9.99");
  });

  it("correctly calculate the subtotal", () => {
    const wrapper = shallowMount(TableOrder, {
      localVue,
      propsData: {
        currentOrder: genTestOrder(),
        refetchFunc() {}
      }
    });

    expect(wrapper.text()).contains("Subtotal $9.99");
  });

  it("correctly calculate the tax", () => {
    const wrapper = shallowMount(TableOrder, {
      localVue,
      propsData: {
        currentOrder: genTestOrder(),
        refetchFunc() {}
      }
    });

    expect(wrapper.text()).contains("Tax $0.70");
  });

  it("correctly calculate the total", () => {
    const wrapper = shallowMount(TableOrder, {
      localVue,
      propsData: {
        currentOrder: genTestOrder(),
        refetchFunc() {}
      }
    });

    expect(wrapper.text()).contains("Total $10.69");
    expect(wrapper.text()).contains("Remaining $10.69");
  });

  it("attempts to create a customer code", done => {
    const mutate = jest.fn();
    const wrapper = mount(TableOrder, {
      localVue,
      propsData: {
        currentOrder: genTestOrder(),
        refetchFunc: () => done()
      },
      mocks: {
        $apollo: {
          mutate
        }
      }
    });

    const custCodeDiv = wrapper.find(".customer-code");
    const _ = expect(custCodeDiv.exists()).to.be.true;
    expect(custCodeDiv.text()).eq("Generate Code");

    const genCustCodeBtn = custCodeDiv.find("button");
    genCustCodeBtn.trigger("click");
  });

  it("opens dialog box for payment", async () => {
    const mutate = jest.fn();
    const wrapper = mount(TableOrder, {
      localVue,
      propsData: {
        currentOrder: genTestOrder(),
        refetchFunc() {}
      },
      mocks: {
        $apollo: {
          mutate
        }
      }
    });

    expect(wrapper.vm.$data.dialogIsOpen).to.equal(false);

    const makePaymentBtn = wrapper.findAll("button").at(1);
    makePaymentBtn.trigger("click");
    await Vue.nextTick();

    expect(wrapper.vm.$data.dialogIsOpen).to.equal(true);
  });

  it("attempts to make a payment", async done => {
    const mutate = jest.fn();
    const wrapper = mount(TableOrder, {
      localVue,
      propsData: {
        currentOrder: genTestOrder(),
        refetchFunc: () => done()
      },
      mocks: {
        $apollo: {
          mutate
        }
      }
    });

    const makePaymentBtn = wrapper.findAll("button").at(1);
    makePaymentBtn.trigger("click");
    await Vue.nextTick();

    wrapper.find(".dialog-ok-btn").trigger("click");
  });
});
