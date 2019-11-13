// No clue why this is needed but tests fail without it
import "regenerator-runtime/runtime";

import { createLocalVue } from "@vue/test-utils";

export const localVue = createLocalVue();
localVue.directive("focus", {
  inserted(el) {
    el.focus();
  }
});
