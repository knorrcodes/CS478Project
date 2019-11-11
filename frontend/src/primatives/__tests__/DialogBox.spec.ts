import { localVue } from "@t/setup-test";
import { shallowMount } from "@vue/test-utils";
import { expect } from "chai";
import Vue from "vue";
import DialogBox from "../DialogBox.vue";

describe("DialogBox", () => {
  it("correctly sets the prompt message", () => {
    const wrapper = shallowMount(DialogBox, {
      localVue,
      propsData: {
        prompt: "Hello",
        okHandler() {},
        cancelHandler() {}
      }
    });

    const promptDir = wrapper.find(".dialog-prompt");
    expect(promptDir.text()).eq("Hello");
  });

  it("correctly updates the input value", async () => {
    const wrapper = shallowMount(DialogBox, {
      localVue,
      propsData: {
        prompt: "Hello",
        okHandler() {},
        cancelHandler() {}
      }
    });

    const input = wrapper.find("input");
    expect(input.element.textContent).eq("");

    input.setValue("10.50");
    await Vue.nextTick();
    expect(wrapper.vm.$data.value).eq("10.50");
  });

  it("correctly triggers okHandler", async done => {
    const okHandler = (data: string) => {
      expect(data).eq("10.50");
      done();
    };

    const wrapper = shallowMount(DialogBox, {
      localVue,
      propsData: {
        prompt: "Hello",
        okHandler,
        cancelHandler() {}
      }
    });

    const input = wrapper.find("input");
    input.setValue("10.50");
    await Vue.nextTick();

    const okBtn = wrapper.find("button");
    okBtn.trigger("click");
  });

  it("correctly triggers cancelHandler", done => {
    const cancelHandler = () => {
      done();
    };

    const wrapper = shallowMount(DialogBox, {
      localVue,
      propsData: {
        prompt: "Hello",
        okHandler() {},
        cancelHandler
      }
    });

    const input = wrapper.find("input");
    input.setValue("10.50");

    const cancelBtn = wrapper.findAll("button").at(1);
    cancelBtn.trigger("click");
  });
});
