/*
 * Copyright 2007-2017 Charles du Jeu - Abstrium SAS <team (at) pyd.io>
 * This file is part of Pydio.
 *
 * Pydio is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Pydio is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with Pydio.  If not, see <http://www.gnu.org/licenses/>.
 *
 * The latest code can be found at <https://pydio.com>.
 */

'use strict';

exports.__esModule = true;

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { 'default': obj }; }

function _interopRequireWildcard(obj) { if (obj && obj.__esModule) { return obj; } else { var newObj = {}; if (obj != null) { for (var key in obj) { if (Object.prototype.hasOwnProperty.call(obj, key)) newObj[key] = obj[key]; } } newObj['default'] = obj; return newObj; } }

var _urls = require('../urls');

var _controls = require('./controls');

var ResolutionControls = _interopRequireWildcard(_controls);

var _actions = require('./actions');

var ResolutionActions = _interopRequireWildcard(_actions);

var _resolution = require('./resolution');

var _resolution2 = _interopRequireDefault(_resolution);

var ResolutionURLProvider = _urls.URLProvider(["hi", "lo"]);

exports.ResolutionURLProvider = ResolutionURLProvider;
exports.ResolutionActions = ResolutionActions;
exports.ResolutionControls = ResolutionControls;
exports.withResolution = _resolution2['default'];
exports.withResolutionControls = _controls.withResolutionControls;
