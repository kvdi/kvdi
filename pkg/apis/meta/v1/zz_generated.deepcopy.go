// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIAction) DeepCopyInto(out *APIAction) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIAction.
func (in *APIAction) DeepCopy() *APIAction {
	if in == nil {
		return nil
	}
	out := new(APIAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthResult) DeepCopyInto(out *AuthResult) {
	*out = *in
	if in.User != nil {
		in, out := &in.User, &out.User
		*out = new(VDIUser)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthResult.
func (in *AuthResult) DeepCopy() *AuthResult {
	if in == nil {
		return nil
	}
	out := new(AuthResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthorizeRequest) DeepCopyInto(out *AuthorizeRequest) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthorizeRequest.
func (in *AuthorizeRequest) DeepCopy() *AuthorizeRequest {
	if in == nil {
		return nil
	}
	out := new(AuthorizeRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionStatus) DeepCopyInto(out *ConnectionStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionStatus.
func (in *ConnectionStatus) DeepCopy() *ConnectionStatus {
	if in == nil {
		return nil
	}
	out := new(ConnectionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CreateRoleRequest) DeepCopyInto(out *CreateRoleRequest) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]Rule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CreateRoleRequest.
func (in *CreateRoleRequest) DeepCopy() *CreateRoleRequest {
	if in == nil {
		return nil
	}
	out := new(CreateRoleRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CreateSessionRequest) DeepCopyInto(out *CreateSessionRequest) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CreateSessionRequest.
func (in *CreateSessionRequest) DeepCopy() *CreateSessionRequest {
	if in == nil {
		return nil
	}
	out := new(CreateSessionRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CreateUserRequest) DeepCopyInto(out *CreateUserRequest) {
	*out = *in
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CreateUserRequest.
func (in *CreateUserRequest) DeepCopy() *CreateUserRequest {
	if in == nil {
		return nil
	}
	out := new(CreateUserRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesktopSession) DeepCopyInto(out *DesktopSession) {
	*out = *in
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(DesktopSessionStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesktopSession.
func (in *DesktopSession) DeepCopy() *DesktopSession {
	if in == nil {
		return nil
	}
	out := new(DesktopSession)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesktopSessionStatus) DeepCopyInto(out *DesktopSessionStatus) {
	*out = *in
	if in.Display != nil {
		in, out := &in.Display, &out.Display
		*out = new(ConnectionStatus)
		**out = **in
	}
	if in.Audio != nil {
		in, out := &in.Audio, &out.Audio
		*out = new(ConnectionStatus)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesktopSessionStatus.
func (in *DesktopSessionStatus) DeepCopy() *DesktopSessionStatus {
	if in == nil {
		return nil
	}
	out := new(DesktopSessionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesktopSessionsResponse) DeepCopyInto(out *DesktopSessionsResponse) {
	*out = *in
	if in.Sessions != nil {
		in, out := &in.Sessions, &out.Sessions
		*out = make([]*DesktopSession, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(DesktopSession)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesktopSessionsResponse.
func (in *DesktopSessionsResponse) DeepCopy() *DesktopSessionsResponse {
	if in == nil {
		return nil
	}
	out := new(DesktopSessionsResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileStat) DeepCopyInto(out *FileStat) {
	*out = *in
	if in.Contents != nil {
		in, out := &in.Contents, &out.Contents
		*out = make([]*FileStat, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(FileStat)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileStat.
func (in *FileStat) DeepCopy() *FileStat {
	if in == nil {
		return nil
	}
	out := new(FileStat)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JWTClaims) DeepCopyInto(out *JWTClaims) {
	*out = *in
	if in.User != nil {
		in, out := &in.User, &out.User
		*out = new(VDIUser)
		(*in).DeepCopyInto(*out)
	}
	out.StandardClaims = in.StandardClaims
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JWTClaims.
func (in *JWTClaims) DeepCopy() *JWTClaims {
	if in == nil {
		return nil
	}
	out := new(JWTClaims)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MFAResponse) DeepCopyInto(out *MFAResponse) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MFAResponse.
func (in *MFAResponse) DeepCopy() *MFAResponse {
	if in == nil {
		return nil
	}
	out := new(MFAResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rule) DeepCopyInto(out *Rule) {
	*out = *in
	if in.Verbs != nil {
		in, out := &in.Verbs, &out.Verbs
		*out = make([]Verb, len(*in))
		copy(*out, *in)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]Resource, len(*in))
		copy(*out, *in)
	}
	if in.ResourcePatterns != nil {
		in, out := &in.ResourcePatterns, &out.ResourcePatterns
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rule.
func (in *Rule) DeepCopy() *Rule {
	if in == nil {
		return nil
	}
	out := new(Rule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SessionResponse) DeepCopyInto(out *SessionResponse) {
	*out = *in
	if in.User != nil {
		in, out := &in.User, &out.User
		*out = new(VDIUser)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SessionResponse.
func (in *SessionResponse) DeepCopy() *SessionResponse {
	if in == nil {
		return nil
	}
	out := new(SessionResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatDesktopFileResponse) DeepCopyInto(out *StatDesktopFileResponse) {
	*out = *in
	if in.Stat != nil {
		in, out := &in.Stat, &out.Stat
		*out = new(FileStat)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatDesktopFileResponse.
func (in *StatDesktopFileResponse) DeepCopy() *StatDesktopFileResponse {
	if in == nil {
		return nil
	}
	out := new(StatDesktopFileResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdateMFARequest) DeepCopyInto(out *UpdateMFARequest) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdateMFARequest.
func (in *UpdateMFARequest) DeepCopy() *UpdateMFARequest {
	if in == nil {
		return nil
	}
	out := new(UpdateMFARequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdateRoleRequest) DeepCopyInto(out *UpdateRoleRequest) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]Rule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdateRoleRequest.
func (in *UpdateRoleRequest) DeepCopy() *UpdateRoleRequest {
	if in == nil {
		return nil
	}
	out := new(UpdateRoleRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdateUserRequest) DeepCopyInto(out *UpdateUserRequest) {
	*out = *in
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdateUserRequest.
func (in *UpdateUserRequest) DeepCopy() *UpdateUserRequest {
	if in == nil {
		return nil
	}
	out := new(UpdateUserRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserMFAStatus) DeepCopyInto(out *UserMFAStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserMFAStatus.
func (in *UserMFAStatus) DeepCopy() *UserMFAStatus {
	if in == nil {
		return nil
	}
	out := new(UserMFAStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VDIUser) DeepCopyInto(out *VDIUser) {
	*out = *in
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]*VDIUserRole, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(VDIUserRole)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.MFA != nil {
		in, out := &in.MFA, &out.MFA
		*out = new(UserMFAStatus)
		**out = **in
	}
	if in.Sessions != nil {
		in, out := &in.Sessions, &out.Sessions
		*out = make([]*DesktopSession, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(DesktopSession)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VDIUser.
func (in *VDIUser) DeepCopy() *VDIUser {
	if in == nil {
		return nil
	}
	out := new(VDIUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VDIUserRole) DeepCopyInto(out *VDIUserRole) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]Rule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VDIUserRole.
func (in *VDIUserRole) DeepCopy() *VDIUserRole {
	if in == nil {
		return nil
	}
	out := new(VDIUserRole)
	in.DeepCopyInto(out)
	return out
}
