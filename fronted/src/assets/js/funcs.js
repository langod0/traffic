import {computed, ref} from "vue";
import {useTransition} from "@vueuse/core";

function pad(num) {
    return num.toString().padStart(2, "0");
}
export function formatDateTime(date) {
    const year = date.getFullYear();
    const month = date.getMonth() + 1;
    const day = date.getDate();
    return `${year}-${pad(month)}-${pad(day)}`;
}
export function useTransitionWithUnit(initialValue, unit, options = {}) {
    const pureValue = initialValue
    const transitionValue = useTransition(pureValue, options)
    const valueWithUnit = computed(() => `${Math.round(transitionValue.value)}${unit}`)

    return valueWithUnit
}

