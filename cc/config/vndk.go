// Copyright 2019 Google Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

// List of VNDK libraries that have different core variant and vendor variant.
// For these libraries, the vendor variants must be installed even if the device
// has VndkUseCoreVariant set.
// Note that AIDL-generated modules must use vendor variants by default.
var VndkMustUseVendorVariantList = []string{
	"android.hardware.nfc@1.2",
	"libbinder",
	"libcrypto",
	"libexpat",
	"libgatekeeper",
	"libgui",
	"libhidlcache",
	"libkeymaster_messages",
	"libkeymaster_portable",
	"libmedia_omx",
	"libpuresoftkeymasterdevice",
	"libselinux",
	"libsoftkeymasterdevice",
	"libsqlite",
	"libssl",
	"libstagefright_bufferpool@2.0",
	"libstagefright_bufferqueue_helper",
	"libstagefright_foundation",
	"libstagefright_omx",
	"libstagefright_omx_utils",
	"libstagefright_xmlparser",
	"libui",
	"libxml2",
	"libmedia_helper",//Remove it from the list once the workaround patch is cleared in S
}
