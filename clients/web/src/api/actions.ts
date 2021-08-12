export function authorize(providerId: string): void {
  window.location.href = `/api/v1/auth/${providerId}/authorize`;
}

export function signOut(): void {
  window.location.href = "/api/v1/auth/signout";
}
