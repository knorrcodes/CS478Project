import { localVue } from "@t/setup-test";
import { shallowMount } from "@vue/test-utils";
import { expect } from "chai";
import Empty from "../Empty.vue";

describe("Empty", () => {
  it("correctly shows slot data", () => {
    const testMsg = "Hello Empty";

    const wrapper = shallowMount(Empty, {
      localVue,
      slots: {
        default: `<span>${testMsg}</span>`
      }
    });

    expect(wrapper.text()).contains(testMsg);
  });

  it("show logout button", () => {
    const wrapper = shallowMount(Empty, {
      localVue
    });

    const logoutBtn = wrapper.find(".logout-btn");
    const _ = expect(logoutBtn.exists()).to.be.true;
  });
});
