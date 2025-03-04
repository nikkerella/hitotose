export interface GameData {
    id: string,
    title: string,
    genre: string,
    platform: string,
    developer: string,
    publisher: string,
    status: string,
    played_time: number,
    time_to_beat: number,
    rating: string,
    ranking: number,
}

export interface StopwatchData {
    game_id: string,
    game_title: string,
    start_time: string,
    end_time: string,
    duration: number,
}

export interface BadgeData {
    played: number,
    playing: number,
    to_play: number,

    all_platform: number,
    pc: number,
    playstation: number,
    nintendo_switch: number,
    xbox: number,
    mobile: number,
}