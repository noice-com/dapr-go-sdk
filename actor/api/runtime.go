/*
Copyright 2021 The Dapr Authors
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package api

import "github.com/dapr/go-sdk/actor/config"

type ActorRuntimeConfig struct {
	RegisteredActorTypes   []string                 `json:"entities"`
	ActorIdleTimeout       string                   `json:"actorIdleTimeout"`
	ActorScanInterval      string                   `json:"actorScanInterval"`
	DrainOngingCallTimeout string                   `json:"drainOngoingCallTimeout"`
	DrainBalancedActors    bool                     `json:"drainRebalancedActors"`
	Reentrancy             *config.ReentrancyConfig `json:"reentrancy"`
}
