function! Generate_uuid() abort
  let l:uuid = system('uuid')
  execute "normal! a".uuid."\<esc>"
endfunction

nnoremap <silent><C-S-u> :call Generate_uuid()<CR>

command! -count -nargs=* UUID call Generate_uuid()
5f13f8ba-6e0c-43c2-9fce-60c99fea0c91
