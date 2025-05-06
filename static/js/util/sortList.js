export function sortClubsByStartDateDesc(clubList) {
    return clubList.sort((a, b) => new Date(b.start_at) - new Date(a.start_at));
}