# Command Series engine

This branch adds audio effects for unit movement and combat, a sound preview tool, and a few bug fixes.

An engine for playing [Command Series](https://www.mobygames.com/game-group/microprose-command-series-games) games ([Crusade in Europe](https://www.mobygames.com/game/crusade-in-europe/), [Decision in the Desert](https://www.mobygames.com/game/decision-in-the-desert/), [Conflict in Vietnam](https://www.mobygames.com/game/conflict-in-vietnam/)) developed by Sid Meier in the mid-eighties and published by MicroProse.

# Using

Obtain an ATR image of Atari version of one of the games and run `$ command_series <diskimage.atr>`.

## Sound preview and tuning

The game uses embedded MP3 recordings for movement and combat:

- `audio/assets/engine.mp3` plays when units move; playback is capped at one second.
- `audio/assets/gunfire.mp3` plays when units attack.
- `audio/assets/nuclear.mp3` plays for explosions and strikes; playback uses the full recording.

Preview the exact sounds used by the game without launching it:

```sh
go run ./tools/sound_preview -sound engine
go run ./tools/sound_preview -sound gunfire
go run ./tools/sound_preview -sound explosion
go run ./tools/sound_preview -sound click -repeat 8
```

Build the preview as a standalone executable if preferred:

```sh
go build -o sound_preview ./tools/sound_preview
./sound_preview -sound explosion
```

The sound names are `engine`, `alert`, `explosion`, `gunfire` (or `shooting`), and `click`. Alert playback is disabled in the game. `engine`, `gunfire`, and `explosion` play the embedded MP3 assets, so replace the corresponding file to change one of those sounds, then rebuild the game:

```sh
go build -o command_series ./cmd
```

Alert and keyboard-click playback are disabled in the game. Alert is also muted in the preview; click remains available there. Engine playback is capped at one second; explosion plays its full recording. Both use 16-bit stereo.

## Tests

Run `go test -race ./...`. Disk-image regression tests use ATRs in the project root.
Set `COMMAND_SERIES_TEST_DATA` to use a different directory. Missing images in the
default directory are reported as skipped tests. The mobile target is compiled
only for Android/iOS and still requires its own `mobile/crusade.atr` asset.

# Missing features

* Bug fixes ~~, many bug-fixes~~
* ~~Save/load~~
* Music (basic movement, report, and combat effects are implemented)
* Intro and ending
