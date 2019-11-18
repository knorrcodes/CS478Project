import { localVue } from "@t/setup-test";
import { mount, RouterLinkStub } from "@vue/test-utils";
import { expect } from "chai";
import RouterLinkStyled from "../RouterLinkStyled.vue";

describe("RouterLinkStyled", () => {
  it("correctly sets link text", () => {
    const wrapper = mount(RouterLinkStyled, {
      localVue,
      stubs: {
        RouterLink: RouterLinkStub
      },
      propsData: {
        value: "Hello",
        to: ""
      }
    });

    const link = wrapper.find("div");
    expect(link.text()).eq("Hello");
  });

  it("correctly sets link destination", () => {
    const wrapper = mount(RouterLinkStyled, {
      localVue,
      stubs: {
        RouterLink: RouterLinkStub
      },
      propsData: {
        value: "Hello",
        to: "/home"
      }
    });

    expect(wrapper.find(RouterLinkStub).props().to).to.eq("/home");
  });
});
