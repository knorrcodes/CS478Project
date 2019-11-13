import { localVue } from "@t/setup-test";
import { mount, shallowMount } from "@vue/test-utils";
import { expect } from "chai";

import StartTableOrder from "../StartTableOrder.vue";
import ButtonStyled from "@/primatives/ButtonStyled.vue";

describe("StartTableOrder", () => {
  it("correctly shows button value", () => {
    const wrapper = mount(StartTableOrder, {
      localVue,
      propsData: {
        startOrder() {}
      }
    });

    const buttonDiv = wrapper.findAll("div").at(1);
    expect(buttonDiv.text()).eq("Start Order");
  });

  it("correctly fires click event", done => {
    const startOrder = () => done();
    const wrapper = mount(StartTableOrder, {
      localVue,
      propsData: {
        startOrder
      }
    });

    const btn = wrapper.find("button");
    btn.trigger("click");
  });
});
