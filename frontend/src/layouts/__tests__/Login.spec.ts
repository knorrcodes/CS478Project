import { localVue } from "@t/setup-test";
import { shallowMount } from "@vue/test-utils";
import { expect } from "chai";
import Login from "../Login.vue";

describe("Login", () => {
  it("correctly shows slot data", () => {
    const testMsg = "Hello Login";

    const wrapper = shallowMount(Login, {
      localVue,
      slots: {
        default: `<span>${testMsg}</span>`
      }
    });

    expect(wrapper.text()).contains(testMsg);
  });

  it("should not show the logout button", () => {
    const wrapper = shallowMount(Login, {
      localVue
    });

    const logoutBtn = wrapper.find(".logout-btn");
    const _ = expect(logoutBtn.exists()).to.be.false;
  });
});
