package v3public

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ADFSProvider) DeepCopyInto(out *ADFSProvider) {
	*out = *in
	in.SamlProvider.DeepCopyInto(&out.SamlProvider)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ADFSProvider.
func (in *ADFSProvider) DeepCopy() *ADFSProvider {
	if in == nil {
		return nil
	}
	out := new(ADFSProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ADFSProvider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveDirectoryProvider) DeepCopyInto(out *ActiveDirectoryProvider) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.AuthProvider.DeepCopyInto(&out.AuthProvider)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveDirectoryProvider.
func (in *ActiveDirectoryProvider) DeepCopy() *ActiveDirectoryProvider {
	if in == nil {
		return nil
	}
	out := new(ActiveDirectoryProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ActiveDirectoryProvider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveDirectoryProviderList) DeepCopyInto(out *ActiveDirectoryProviderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ActiveDirectoryProvider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveDirectoryProviderList.
func (in *ActiveDirectoryProviderList) DeepCopy() *ActiveDirectoryProviderList {
	if in == nil {
		return nil
	}
	out := new(ActiveDirectoryProviderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ActiveDirectoryProviderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthProvider) DeepCopyInto(out *AuthProvider) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthProvider.
func (in *AuthProvider) DeepCopy() *AuthProvider {
	if in == nil {
		return nil
	}
	out := new(AuthProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AuthProvider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthProviderList) DeepCopyInto(out *AuthProviderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AuthProvider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthProviderList.
func (in *AuthProviderList) DeepCopy() *AuthProviderList {
	if in == nil {
		return nil
	}
	out := new(AuthProviderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AuthProviderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureADLogin) DeepCopyInto(out *AzureADLogin) {
	*out = *in
	out.GenericLogin = in.GenericLogin
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureADLogin.
func (in *AzureADLogin) DeepCopy() *AzureADLogin {
	if in == nil {
		return nil
	}
	out := new(AzureADLogin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureADProvider) DeepCopyInto(out *AzureADProvider) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.AuthProvider.DeepCopyInto(&out.AuthProvider)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureADProvider.
func (in *AzureADProvider) DeepCopy() *AzureADProvider {
	if in == nil {
		return nil
	}
	out := new(AzureADProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureADProvider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BasicLogin) DeepCopyInto(out *BasicLogin) {
	*out = *in
	out.GenericLogin = in.GenericLogin
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BasicLogin.
func (in *BasicLogin) DeepCopy() *BasicLogin {
	if in == nil {
		return nil
	}
	out := new(BasicLogin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FreeIpaProvider) DeepCopyInto(out *FreeIpaProvider) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.AuthProvider.DeepCopyInto(&out.AuthProvider)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FreeIpaProvider.
func (in *FreeIpaProvider) DeepCopy() *FreeIpaProvider {
	if in == nil {
		return nil
	}
	out := new(FreeIpaProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FreeIpaProvider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericLogin) DeepCopyInto(out *GenericLogin) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericLogin.
func (in *GenericLogin) DeepCopy() *GenericLogin {
	if in == nil {
		return nil
	}
	out := new(GenericLogin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GithubLogin) DeepCopyInto(out *GithubLogin) {
	*out = *in
	out.GenericLogin = in.GenericLogin
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GithubLogin.
func (in *GithubLogin) DeepCopy() *GithubLogin {
	if in == nil {
		return nil
	}
	out := new(GithubLogin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GithubProvider) DeepCopyInto(out *GithubProvider) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.AuthProvider.DeepCopyInto(&out.AuthProvider)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GithubProvider.
func (in *GithubProvider) DeepCopy() *GithubProvider {
	if in == nil {
		return nil
	}
	out := new(GithubProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GithubProvider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyCloakProvider) DeepCopyInto(out *KeyCloakProvider) {
	*out = *in
	in.SamlProvider.DeepCopyInto(&out.SamlProvider)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyCloakProvider.
func (in *KeyCloakProvider) DeepCopy() *KeyCloakProvider {
	if in == nil {
		return nil
	}
	out := new(KeyCloakProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeyCloakProvider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalProvider) DeepCopyInto(out *LocalProvider) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.AuthProvider.DeepCopyInto(&out.AuthProvider)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalProvider.
func (in *LocalProvider) DeepCopy() *LocalProvider {
	if in == nil {
		return nil
	}
	out := new(LocalProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalProvider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OKTAProvider) DeepCopyInto(out *OKTAProvider) {
	*out = *in
	in.SamlProvider.DeepCopyInto(&out.SamlProvider)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OKTAProvider.
func (in *OKTAProvider) DeepCopy() *OKTAProvider {
	if in == nil {
		return nil
	}
	out := new(OKTAProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OKTAProvider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenLdapProvider) DeepCopyInto(out *OpenLdapProvider) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.AuthProvider.DeepCopyInto(&out.AuthProvider)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenLdapProvider.
func (in *OpenLdapProvider) DeepCopy() *OpenLdapProvider {
	if in == nil {
		return nil
	}
	out := new(OpenLdapProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenLdapProvider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PingProvider) DeepCopyInto(out *PingProvider) {
	*out = *in
	in.SamlProvider.DeepCopyInto(&out.SamlProvider)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PingProvider.
func (in *PingProvider) DeepCopy() *PingProvider {
	if in == nil {
		return nil
	}
	out := new(PingProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PingProvider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SamlLoginInput) DeepCopyInto(out *SamlLoginInput) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SamlLoginInput.
func (in *SamlLoginInput) DeepCopy() *SamlLoginInput {
	if in == nil {
		return nil
	}
	out := new(SamlLoginInput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SamlLoginOutput) DeepCopyInto(out *SamlLoginOutput) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SamlLoginOutput.
func (in *SamlLoginOutput) DeepCopy() *SamlLoginOutput {
	if in == nil {
		return nil
	}
	out := new(SamlLoginOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SamlProvider) DeepCopyInto(out *SamlProvider) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.AuthProvider.DeepCopyInto(&out.AuthProvider)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SamlProvider.
func (in *SamlProvider) DeepCopy() *SamlProvider {
	if in == nil {
		return nil
	}
	out := new(SamlProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SamlProvider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
