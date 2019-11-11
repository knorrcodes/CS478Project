import { localVue } from "@t/setup-test";
import { shallowMount } from "@vue/test-utils";
import { expect } from "chai";
import ButtonStyled from "../ButtonStyled.vue";

describe("ButtonStyled", () => {
  it("correctly sets button value", () => {
    const wrapper = shallowMount(ButtonStyled, {
      localVue,
      propsData: {
        value: "Hello",
        clickHandler() {}
      }
    });

    const btn = wrapper.find("button");
    expect(btn.text()).eq("Hello");
  });

  it("correctly triggers clickHandler", done => {
    const wrapper = shallowMount(ButtonStyled, {
      localVue,
      propsData: {
        value: "Hello",
        clickHandler: () => done()
      }
    });

    const btn = wrapper.find("button");
    btn.trigger("click");
  });
});
