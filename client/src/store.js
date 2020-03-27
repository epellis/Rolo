import { writable } from 'svelte/store';

function createUserStore() {
    const defaultUser = { isLoggedIn: false }

    const { subscribe, set, update } = writable(defaultUser);

    return {
        subscribe,
        logIn: (userData) => {
            userData.isLoggedIn = true;
            set(userData);
        },
        logOut: () => set(defaultUser)
    };
}

export const userStore = createUserStore();
